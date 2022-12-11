package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {
	conn := NewConn()
	err := conn.login()
	assert.Nil(t, err)
}
