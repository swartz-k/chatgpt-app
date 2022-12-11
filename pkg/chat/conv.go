package chat

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/swartz-k/chatgpt-app/entity/valobj"
	"github.com/swartz-k/chatgpt-app/pkg/log"
	expire "github.com/swartz-k/chatgpt-app/pkg/utils/map"
	"net/http"
	"strings"
	"time"
)

type API struct {
	sessionToken      string
	markdown          bool
	apiBaseUrl        string
	backendApiBaseUrl string
	userAgent         string
	tokenCache        expire.ExpireMap
}

func NewApi(session *string) *API {
	return &API{
		sessionToken:      *session,
		markdown:          true,
		apiBaseUrl:        "https://chat.openai.com/api",
		backendApiBaseUrl: "https://chat.openai.com/backend-api",
		userAgent:         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36",
		tokenCache:        expire.ExpireMap{},
	}
}

func (a *API) refreshToken() (*string, error) {
	expireRecord := a.tokenCache.Get(a.sessionToken)
	if expireRecord != nil {
		return &expireRecord.Value, nil
	}
	token, err := getTokenBySession(a.sessionToken)
	if err != nil {
		return nil, errors.Wrap(err, "get access token with session")
	}
	//token := login()
	a.tokenCache.Set(a.sessionToken, &expire.ExpireRecord{Expire: time.Now().Add(time.Second * 10), Value: token.Value})
	return &token.Value, nil
}

func (a *API) SendMessage(msg *valobj.Message) (*valobj.ConversationResponse, error) {
	token, err := a.refreshToken()
	if err != nil {
		return nil, err
	}

	if msg.ParentId == "" {
		msg.ParentId = uuid.NewString()
	}
	body := valobj.ConversationJSONBody{
		Action: "next",
		Messages: []valobj.Prompt{{
			Id:      uuid.NewString(),
			Role:    valobj.RoleUser,
			Content: valobj.PromptContent{ContentType: valobj.ContentTypeText, Parts: []string{msg.Message}},
		}},
		Model:           "text-davinci-002-render",
		ParentMessageId: msg.ParentId,
	}
	if msg.ConvId != "" {
		body.ConversationId = msg.ConvId
	}
	log.V(100).Info("send payload %+v", body)
	url := fmt.Sprintf("%s/conversation", a.backendApiBaseUrl)

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", a.userAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r *valobj.ConversationResponse
	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(ScanTwoLines)
	for scanner.Scan() {
		content := scanner.Bytes()
		tmp, err := valobj.GetConversationResponse(content)
		if err != nil {
			log.V(101).Info("scan %s unmarshal err %+v", content, err)
		}
		if tmp != nil {
			r = tmp
		}
	}

	return r, nil
}

func (a *API) WatchMessage(msg *valobj.Message) chan *valobj.ConversationResponse {
	ch := make(chan *valobj.ConversationResponse, 3)
	token, err := a.refreshToken()
	if err != nil {
		log.V(100).Info("refresh token failed %+v", err)
		return nil
	}

	if msg.ParentId == "" {
		msg.ParentId = uuid.NewString()
	}
	body := valobj.ConversationJSONBody{
		Action: "next",
		Messages: []valobj.Prompt{{
			Id:      uuid.NewString(),
			Role:    valobj.RoleUser,
			Content: valobj.PromptContent{ContentType: valobj.ContentTypeText, Parts: []string{msg.Message}},
		}},
		Model:           "text-davinci-002-render",
		ParentMessageId: msg.ParentId,
	}
	if msg.ConvId != "" {
		body.ConversationId = msg.ConvId
	}
	log.V(100).Info("send payload %+v", body)
	url := fmt.Sprintf("%s/conversation", a.backendApiBaseUrl)

	payload, err := json.Marshal(body)
	if err != nil {
		log.V(100).Info("payload marshal  failed %+v", err)
		return nil
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		log.V(100).Info("new request failed %+v", err)
		return nil
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *token))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", a.userAgent)

	go func() {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.V(100).Info("do request failed %+v", err)
			return
		}
		log.V(100).Info("response code %d", resp.StatusCode)
		defer resp.Body.Close()
		scanner := bufio.NewScanner(resp.Body)
		scanner.Split(ScanTwoLines)
		for scanner.Scan() {
			content := scanner.Bytes()
			log.V(100).Info("scanner receive msg %s", content)
			tmp, err := valobj.GetConversationResponse(content)
			if err != nil {
				if strings.Contains(err.Error(), "DONE") {
					ch <- &valobj.ConversationResponse{Error: err}
					return
				}
				log.V(101).Info("scan %s unmarshal err %+v", content, err)
			}

			ch <- tmp
		}
	}()
	return ch
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanTwoLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexAny(data, "\n\n"); i >= 0 {
		return i + 1, dropCR(data[0:i]), nil
	}
	if atEOF {
		return len(data), dropCR(data), nil
	}
	return 0, nil, nil
}
