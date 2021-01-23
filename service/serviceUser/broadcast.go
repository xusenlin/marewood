package serviceUser

import (
	"MareWood/models"
	"github.com/gorilla/websocket"
)

func JoinWsClientsByToken(token string, conn *websocket.Conn) error {
	claims, err := ParseToken(token)
	if err != nil {
		return err
	}

	if conn,ok := models.WsClients[claims.ID]; ok {
		_ = conn.WsConn.Close()
	}
	models.WsClients[claims.ID] = models.ConnUser{Username: claims.Username, WsConn: conn}

	models.Broadcast <- models.Message{
		Type:            models.MsgTypeInfo,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		Message:         "“" + claims.Username + "” Joined the system",
	}

	return nil
}
