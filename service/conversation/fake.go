package conversation

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/swartz-k/chatgpt-app/entity/valobj"
)

type fakeConversation struct {
	conversationId  string
	parentMessageId string
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
