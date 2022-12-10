package service

import (
	"github.com/swartz-k/chatgpt-app/service/conversation"
	"sync"
)

var (
	conversationOnce sync.Once
	convMapper       map[string]conversation.Conversation
)

func GetConversationService(token string) conversation.Conversation {
	conversationOnce.Do(func() {
		if convMapper == nil {
			convMapper = map[string]conversation.Conversation{}
		}
	})
	if conv, ok := convMapper[token]; ok {
		return conv
	}
	conv := conversation.NewConversation()
	convMapper[token] = conv
	return conv
}
