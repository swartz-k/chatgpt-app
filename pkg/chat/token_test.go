package chat

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func Test_getChat(t *testing.T) {
	session := "session"
	token, err := getTokenBySession(session)
	assert.Nil(t, err)
	t.Logf("session %v", token)
}

func Test_listModules(t *testing.T) {
	key := "sk-abcdef"
	req, err := http.NewRequest(http.MethodGet, "https://api.openai.com/v1/models", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", key))
	response, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	t.Logf("modules \n %s", content)
}
