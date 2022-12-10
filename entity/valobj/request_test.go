package valobj

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_unmarshalResult(t *testing.T) {
	r := &SessionResult{}
	content := `{"user":{"id":"a","name":"b","email":"c","image":"d","picture":"e","groups":[],"features":[]},"expires":"2023-01-08T23:18:30.506Z","accessToken":"D8O8S_GVQFmRUuw"}`
	err := json.Unmarshal([]byte(content), r)
	assert.Nil(t, err)
	t.Logf("result %+v", r)
}

func TestGetConversationResponse(t *testing.T) {
	data := `data: {"message": {"id": "a", "role": "assistant", "user": null, "create_time": null, "update_time": null, "content": {"content_type": "text", "parts": ["I am a large language model trained by OpenAI."]}, "end_turn": null, "weight": 1.0, "metadata": {}, "recipient": "all"}, "conversation_id": "9038716308dc6", "error": null}`
	c, err := GetConversationResponse([]byte(data))
	assert.Nil(t, err)
	t.Logf("conversation %v", c)
}
