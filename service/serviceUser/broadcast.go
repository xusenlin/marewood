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
	models.WsClients[claims.ID] = models.ConnUser{Username: claims.Username, WsConn: conn}

	models.Broadcast <- models.Message{
		Type:            models.MsgTypeInfo,
		TriggerID:       claims.ID,
		TriggerUsername: claims.Username,
		Message:         "“" + claims.Username + "” Joined the system",
	}

	return nil
}
