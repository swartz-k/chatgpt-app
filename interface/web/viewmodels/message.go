package viewmodels

type Message struct {
	ConvId   string `json:"con_id"`
	ParentId string `json:"par_id"`
	Message  string `json:"message"`
	// for auth
	Token string `json:"tok"`
}
