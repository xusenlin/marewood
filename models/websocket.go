package models

import (
	"fmt"
	"github.com/gorilla/websocket"
)

const (
	MsgTypeInfo    = "info"
	MsgTypeError   = "error"
	MsgTypeSuccess = "success"
	MsgTypeWarning = "warning"
)

const (
	UpdateDataTypeIsNotice = iota
	UpdateDataTypeIsRepoAction
	UpdateDataTypeIsCategoryAction
	UpdateDataTypeIsJobAction
	UpdateDataTypeIsUserAction
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
	UpdateDataType  uint
	Message         string
}

var WsClients = make(map[uint]ConnUser)

var Broadcast = make(chan Message)


func BroadcastMessages() {
	for {
		msg := <-Broadcast
		fmt.Println(WsClients)
		for id, client := range WsClients {
			if !msg.NeedNotifySelf && msg.TriggerID == id {
				//如果不需要通知自己
				continue
			}
			err := client.WsConn.WriteJSON(msg)
			if err != nil {
				client.WsConn.Close()
				delete(WsClients, id)
			}
		}
	}
}
