package chat

import (
	"encoding/json"
	"github.com/swartz-k/chatgpt-app/entity/valobj"
	"github.com/swartz-k/chatgpt-app/pkg/log"
	expire "github.com/swartz-k/chatgpt-app/pkg/utils/map"
	"io"
	"net/http"
	"net/http/cookiejar"
)

var client http.Client

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal("Got error while creating cookie jar %s", err.Error())
	}
	client = http.Client{
		Jar: jar,
	}
}

func getTokenBySession(session string) (*expire.ExpireRecord, error) {
	url := "https://chat.openai.com/api/auth/session"
	agent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	cookie := &http.Cookie{
		Name:  "__Secure-next-auth.session-token",
		Value: session,
	}
	req.AddCookie(cookie)
	req.Header.Add("user-agent", agent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r valobj.SessionResult
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.V(100).Info("content %s", content)
	err = json.Unmarshal(content, &r)
	if err != nil {
		return nil, err
	}
	log.V(100).Info("result %+v, err %+v", r, err)

	return &expire.ExpireRecord{Expire: r.Expires, Value: r.AccessToken}, nil
}
