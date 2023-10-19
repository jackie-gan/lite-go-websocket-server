package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{}

func wsHandler(w http.ResponseWriter, req *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, req, nil)
	if err != nil {
		panic(err)
	}

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func main() {
	http.HandleFunc("/", wsHandler)
	fmt.Println("socket server started on port 9099...")
	http.ListenAndServe(":9099", nil)
}
