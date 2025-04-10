package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Upgrader to upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // allow all origins
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected:", r.RemoteAddr)

	// Goroutine to send messages every 2 seconds
	go func() {
		for {
			msg := fmt.Sprintf("Server time: %s", time.Now().Format(time.RFC3339))
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				log.Println("Write error:", err)
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()

	// Read messages from client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received from client: %s\n", message)
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	port := "8080"
	log.Printf("WebSocket server started on ws://localhost:%s/ws", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
