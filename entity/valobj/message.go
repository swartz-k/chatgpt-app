package valobj

type Message struct {
	ConvId   string `json:"convId"`
	ParentId string `json:"parentId"`
	Message  string `json:"message"`
}
