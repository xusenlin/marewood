package controller

import (
	"MareWood/models"
	"MareWood/service/serviceUser"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketMsg(c *gin.Context) {
	token := c.Query("token")
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil || token == "" {
		return
	}
	err = serviceUser.JoinWsClientsByToken(token, ws)
	if err != nil {
		_ = ws.WriteJSON(models.Message{
			Type:    models.MsgTypeError,
			Message: err.Error(),
		})
		_ = ws.Close()
	}
}
