package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/swartz-k/chatgpt-app/entity/valobj"
	"github.com/swartz-k/chatgpt-app/interface/web/viewmodels"
	"github.com/swartz-k/chatgpt-app/pkg/log"
	"github.com/swartz-k/chatgpt-app/service"
	"net/http"
)

func WechatMessage(c *gin.Context) (int, interface{}, error) {
	var r viewmodels.Message
	err := c.Bind(&r)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	var token string
	if r.Token != "" {
		token = r.Token
	} else {
		token = uuid.NewString()
	}

	msg, err := service.GetConversationService(token).Send(
		&valobj.Message{ConvId: r.ConvId, ParentId: r.ParentId, Message: r.Message})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	log.V(100).Info("response %v", msg.Message)
	mv := viewmodels.ToMessageView(msg, token)
	return http.StatusOK, &mv, nil
}

func SocketMessage(c *gin.Context, conn *websocket.Conn) error {
	_, content, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	m := viewmodels.Message{}
	err = json.Unmarshal(content, &m)
	if err != nil {
		log.V(100).Info("unmarshal socket message err %+v", err)
		return err
	}
	defer conn.Close()

	var token string
	if m.Token != "" {
		token = m.Token
	} else {
		token = uuid.NewString()
	}
	err = service.GetConversationService(token).Watch(&valobj.Message{ConvId: m.ConvId, ParentId: m.ParentId, Message: m.Message}, conn)

	log.V(100).Info("write websocket conn,  err %+v", err)

	return nil
}
