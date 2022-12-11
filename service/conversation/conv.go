package conversation

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/swartz-k/chatgpt-app/entity/valobj"
	"github.com/swartz-k/chatgpt-app/pkg/chat"
	"github.com/swartz-k/chatgpt-app/pkg/config"
	"github.com/swartz-k/chatgpt-app/pkg/log"
	"strings"
	"sync"
	"time"
)

type Conversation interface {
	Send(message *valobj.Message) (*valobj.Message, error)
	Watch(message *valobj.Message, conn *websocket.Conn) error
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

	r, err := c.api.SendMessage(message)
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

func (c *chatConversation) Watch(message *valobj.Message, conn *websocket.Conn) error {
	if message == nil {
		return fmt.Errorf("cannot send nil message")
	}
	c.conversationId = message.ConvId
	c.parentMessageId = message.ParentId

	ch := c.api.WatchMessage(message)
	if ch == nil {
		return fmt.Errorf("message without response")
	}
	t := time.NewTimer(time.Minute * 4)
	for {
		select {
		case <-t.C:
			err := conn.WriteJSON(&valobj.Message{ConvId: c.conversationId, ParentId: c.parentMessageId, Message: "server timeout"})
			if err != nil {
				return err
			}
		case r := <-ch:
			if r == nil {
				continue
			}
			if r.Error != nil {
				log.V(100).Info("receive error %v", r.Error)
				if strings.Contains(r.Error.Error(), "DONE") {
					return nil
				}
			}
			c.conversationId = r.ConversationId
			c.parentMessageId = r.Message.Id
			msg := r.GetMessage()
			if msg == nil {
				log.V(100).Info("get message from %+v is nil", r)
				continue
			}
			err := conn.WriteJSON(&valobj.Message{ConvId: r.ConversationId, ParentId: r.Message.Id, Message: *msg})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
