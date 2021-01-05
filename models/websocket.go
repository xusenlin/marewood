package models

import (
	"github.com/gorilla/websocket"
)

const (
	MsgTypeInfo    = "info"
	MsgTypeError   = "error"
	MsgTypeSuccess = "success"
	MsgTypeWarning = "warning"
)

type ConnUser struct {
	Username string
	WsConn   *websocket.Conn
}

type Message struct {
	Type            string
	TriggerID       uint
	TriggerUsername string
	NeedNotifySelf  bool
	Message         string
	OnlineUsers     int
}

var WsClients = make(map[uint]ConnUser)

var Broadcast = make(chan Message)

func BroadcastMessages() {
	for {
		msg := <-Broadcast
		msg.OnlineUsers = len(WsClients)
		for id, client := range WsClients {
			if !msg.NeedNotifySelf && msg.TriggerID == id {
				//如果不需要通知自己
				continue
			}
			err := client.WsConn.WriteJSON(msg)
			if err != nil {
				_ = client.WsConn.Close()
				delete(WsClients, id)
			}
		}
	}
}
