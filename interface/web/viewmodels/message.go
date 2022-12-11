package viewmodels

import "github.com/swartz-k/chatgpt-app/entity/valobj"

type Message struct {
	ConvId   string `json:"con_id"`
	ParentId string `json:"par_id"`
	Message  string `json:"message"`
	// for auth
	Token string `json:"tok"`
}

func ToMessageView(m *valobj.Message, token string) *Message {
	return &Message{
		ConvId:   m.ConvId,
		ParentId: m.ParentId,
		Message:  m.Message,
		Token:    token,
	}
}
