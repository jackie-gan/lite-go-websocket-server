package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const serverPort = "9099"

var wsUpgrader = websocket.Upgrader{}

func wsHandler(ctx *gin.Context) {
	conn, err := wsUpgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		panic(err)
	}

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Printf("receive: %s\n", string(msg))

		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func logHandler(ctx *gin.Context) {
	fmt.Printf("request from %s, url: %s\n", ctx.Request.RemoteAddr, ctx.Request.URL)
	ctx.Next()
}

func main() {
	server := gin.Default()

	server.Use(logHandler, wsHandler)

	fmt.Println("socket server started on port 9099...")
	server.Run(":" + serverPort)
}
