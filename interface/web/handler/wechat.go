package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	mv := viewmodels.Message{ConvId: msg.ConvId, ParentId: msg.ParentId, Message: msg.Message, Token: token}
	return http.StatusOK, &mv, nil
}
