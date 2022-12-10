package valobj

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/swartz-k/chatgpt-app/pkg/log"
	"time"
)

type ContentType string
type Role string

var (
	ContentTypeText ContentType = "text"
	RoleUser        Role        = "user"
	RoleAssistant   Role        = "assistant"
)

type User struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Image    string   `json:"image"`
	Picture  string   `json:"picture"`
	Groups   []string `json:"groups"`
	Features []string `json:"features"`
}

type SessionResult struct {
	User        User      `json:"user"`
	Expires     time.Time `json:"expires"`
	AccessToken string    `json:"accessToken"`
	Error       string    `json:"error,omitempty"`
}

type PromptContent struct {
	ContentType ContentType `json:"content_type"`
	Parts       []string    `json:"parts"`
}

type Prompt struct {
	Content PromptContent `json:"content"`
	Id      string        `json:"id"`
	Role    Role          `json:"role"`
}

type ConversationJSONBody struct {
	Action          string   `json:"action"`
	ConversationId  string   `json:"conversation_id,omitempty"`
	Messages        []Prompt `json:"messages"`
	Model           string   `json:"model"`
	ParentMessageId string   `json:"parent_message_id"`
}

type ConversationResponse struct {
	Message        Prompt `json:"message"`
	ConversationId string `json:"conversation_id"`
	Error          error  `json:"error,omitempty"`
}

func GetConversationResponse(content []byte) (*ConversationResponse, error) {
	if len(content) < 2 {
		return nil, fmt.Errorf("content %s is invalid", content)
	}
	result := bytes.Split(content, []byte("data: "))
	var c ConversationResponse
	if len(result) == 0 {
		return nil, fmt.Errorf("cannot unmarshal %s", content)
	}
	var maxIndex int
	var max int
	for i := 0; i < len(result); i++ {
		if len(result[i]) > max {
			max = len(result[i])
			maxIndex = i
		}
	}
	err := json.Unmarshal(result[maxIndex], &c)
	if err != nil {
		log.V(100).Info("cannot unmarshal %s", result[maxIndex])
		return nil, err
	}
	return &c, nil
}

func (r *ConversationResponse) GetMessage() *string {
	if r != nil && r.Message.Content.Parts != nil {
		return &r.Message.Content.Parts[0]
	}
	return nil
}
