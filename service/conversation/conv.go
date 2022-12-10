package conversation

import (
	"fmt"
	"github.com/swartz-k/chatgpt-app/entity/valobj"
	"github.com/swartz-k/chatgpt-app/pkg/chat"
	"github.com/swartz-k/chatgpt-app/pkg/config"
	"sync"
)

type Conversation interface {
	Send(message *valobj.Message) (*valobj.Message, error)
}

var (
	apiOnce sync.Once
	api     *chat.API
)

func getApi() *chat.API {
	apiOnce.Do(func() {
		cfg := config.GetConfig(nil)
		api = chat.NewApi(&cfg.Session)
	})
	return api
}

type chatConversation struct {
	api             *chat.API
	conversationId  string
	parentMessageId string
}

func NewConversation() Conversation {
	conv := chatConversation{
		api: getApi(),
	}
	return &conv
}

func (c *chatConversation) Send(message *valobj.Message) (*valobj.Message, error) {
	if message == nil {
		return nil, fmt.Errorf("cannot send nil message")
	}
	c.conversationId = message.ConvId
	c.parentMessageId = message.ParentId

	r, err := c.api.SenMessage(message)
	if err != nil {
		return nil, err
	}
	c.conversationId = r.ConversationId
	c.parentMessageId = r.Message.Id
	msg := r.GetMessage()
	if msg == nil {
		return nil, fmt.Errorf("get message from %+v is nil", r)
	}

	return &valobj.Message{ConvId: r.ConversationId, ParentId: r.Message.Id, Message: *msg}, nil
}
