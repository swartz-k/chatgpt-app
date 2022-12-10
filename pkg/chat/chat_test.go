package chat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getChat(t *testing.T) {
	session := "session"
	token, err := getTokenBySession(session)
	assert.Nil(t, err)
	t.Logf("session %v", token)
}
