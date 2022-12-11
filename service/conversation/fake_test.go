package conversation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fakeWatch(t *testing.T) {
	conv := NewFakeConversation()
	err := conv.Watch(nil, nil)
	assert.Nil(t, err)
}
