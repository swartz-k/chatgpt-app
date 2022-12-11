package conversation

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/swartz-k/chatgpt-app/entity/valobj"
	"time"
)

type fakeConversation struct {
	conversationId  string
	parentMessageId string
}

func chanMsg() chan *valobj.ConversationResponse {
	ch := make(chan *valobj.ConversationResponse, 10)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			fmt.Printf("send msg \n")
			ch <- &valobj.ConversationResponse{ConversationId: time.Now().String()}
		}
	}()
	return ch
}

func (c *fakeConversation) Watch(message *valobj.Message, conn *websocket.Conn) error {
	ch := chanMsg()
	for {
		select {
		case r := <-ch:
			fmt.Printf("receive msg %+v\n", r)
		}
	}
}

func NewFakeConversation() Conversation {
	conv := fakeConversation{
		conversationId:  uuid.NewString(),
		parentMessageId: uuid.NewString(),
	}
	return &conv
}

func (c *fakeConversation) Send(message *valobj.Message) (*valobj.Message, error) {
	if message == nil {
		return nil, fmt.Errorf("cannot send nil message")
	}
	c.conversationId = message.ConvId
	c.parentMessageId = message.ParentId

	return &valobj.Message{ConvId: c.conversationId, ParentId: c.parentMessageId, Message: message.Message}, nil
}
