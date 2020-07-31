package models

import (
	"fmt"
	"github.com/gorilla/websocket"
)

const (
	MsgTypeIsNotice = iota
	MsgTypeIsRepoBuildSuccess
	MsgTypeIsJobBuildSuccess
)

type ConnUser struct {
	Username string
	WsConn   *websocket.Conn
}

type Message struct {
	Type            uint
	TriggerID       uint
	TriggerUsername string
	Message         string
}

var WsClients = make(map[uint]ConnUser)

var Broadcast = make(chan Message)

func BroadcastMessages() {
	for {
		msg := <-Broadcast
		fmt.Println("=========================")
		fmt.Println(WsClients)
		fmt.Println("=========================")
		for id, client := range WsClients {
			if msg.TriggerID != id{
				err := client.WsConn.WriteJSON(msg)
				if err != nil {
					client.WsConn.Close()
					delete(WsClients, id)
				}
			}
		}
	}
}
