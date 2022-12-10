package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/swartz-k/chatgpt-app/interface/web/handler"
	"github.com/swartz-k/chatgpt-app/pkg/log"
	"net/http"
)

type WsRoute struct {
	Path    string
	Method  string
	Handler func(c *gin.Context, conn *websocket.Conn) error
	Flag    []string
}

type Handler func(c *gin.Context) (int, interface{}, error)

type ApiRoute struct {
	Path    string
	Method  string
	Handler Handler
	Flag    map[string]bool
}

type response struct {
	Msg   interface{} `json:"msg,omitempty"`
	Data  interface{} `json:"data"`
	Error string      `json:"error,omitempty"`
}

// todo: message progress
var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func websocketWrapper(f func(c *gin.Context, conn *websocket.Conn) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Del("Origin")
		conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		defer conn.Close()
		err = f(c, conn)
		if err != nil {
			log.V(9).Info("failed websockt err %+v", err)
			return
		}
	}
}

func Register(route *gin.Engine) {
	// init handler interface
	// register http route
	api := route.Group("/chat/api/v1")

	apiV1WechatRoute := []ApiRoute{
		{Path: "/wechat/message", Method: http.MethodPost, Handler: handler.WechatMessage},
	}
	for _, i := range apiV1WechatRoute {
		h := i
		funcHandler := func(c *gin.Context) {
			code, msg, err := h.Handler(c)
			log.V(1).Info("wechat req: path %s, code %d, err %+v", c.FullPath(), code, err)
			if err != nil || code != http.StatusOK {
				c.JSON(code, response{Msg: "", Error: err.Error(), Data: ""})
			} else {
				c.JSON(code, msg)
			}
		}
		api.Handle(h.Method, h.Path, funcHandler)
	}

	ws := route.Group("/ws/v1")
	var wsV1Route []WsRoute
	for _, w := range wsV1Route {
		ws.Handle(w.Method, w.Path, websocketWrapper(w.Handler))
	}
}
