package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Printf("Connect socket error :%s\n", err.Error())
			return
		}

		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				fmt.Printf("Read input stream error: %s\n", err.Error())
				break
			}
			fmt.Printf("%s send: %s\n", c.RemoteAddr().String(), string(msg))
			if err = c.WriteMessage(mt, msg); err != nil {
				fmt.Printf("Write message error: %s\n", err.Error())
				break
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8001", nil))
}
