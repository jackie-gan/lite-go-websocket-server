package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://49.51.65.244/sandbox", nil)
	if err != nil {
		fmt.Println("Websocket Connect Error: ", err)
		panic(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, msg, err := conn.ReadMessage()
				if err != nil {
					fmt.Println("Websocket Read Message Error: ", err)
					cancel()
					return
				}

				os.Stdout.Write(msg)
			}
		}
	}()

	go func() {
		inputReader := bufio.NewReader(os.Stdin)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				input, err := inputReader.ReadString('\n')
				if err != nil {
					cancel()
				}

				fmt.Println("Send Message: ", input)
				err = conn.WriteMessage(websocket.TextMessage, []byte(input))
				if err != nil {
					fmt.Println("Websocket Send Error: ", err)
					return
				}
			}
		}
	}()

	select {
	case <-ctx.Done():
		return
	}
}
