package chat

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/swartz-k/chatgpt-app/entity/valobj"
	"testing"
)

func Test_uuid(t *testing.T) {
	u := uuid.NewString()
	t.Logf("uuid %s", u)
}

func Test_apiSenMessage(t *testing.T) {
	session := "ey"
	a := NewApi(&session)
	message := &valobj.Message{Message: "Where are you from?"}
	r, err := a.SendMessage(message)
	assert.Nil(t, err)
	msg := r.GetMessage()
	assert.NotNil(t, msg)
	t.Logf("resp %+v", *msg)
}
