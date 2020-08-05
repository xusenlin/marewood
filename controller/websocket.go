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

	claims, err := serviceUser.ParseToken(token)

	if err != nil {
		return
	}
	//defer ws.Close()

	models.WsClients[claims.ID] = models.ConnUser{Username: claims.Username, WsConn: ws}

	msg := models.Message{
		Type:            models.MsgTypeInfo,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		UpdateDataType:  models.UpdateDataTypeIsNotice,
		Message:         "“" + claims.Username + "” 加入了系统",
	}
	models.Broadcast <- msg
	//读取用户信息
	//for {
	//	var msg models.Message
	//	//err := ws.ReadJSON(&msg)
	//	//if err != nil {
	//	//	log.Printf("error: %v", err)
	//	//	delete(clients, ws)
	//	//	break
	//	//}
	//	// Send the newly received message to the broadcast channel
	//	broadcast <- msg
	//}
}
