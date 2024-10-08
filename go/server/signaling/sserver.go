package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	addr = "localhost:7071"
)

// htsupgrader is HTTP server upgrader
// Upgraded to Websocket Connection
var htsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func SocketServer(rw http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP request to a WebSocket connection
	conn, err := htsupgrader.Upgrade(rw, r, nil)
	if err != nil {
		http.Error(rw, "Failed to upgrade to WebSocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}

		log.Printf("Received: %s", message)

		// Echo message back to client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}

func main() {
	// Register the WebSocket handler at the /ws endpoint
	http.HandleFunc("/ws", SocketServer)

	log.Printf("Server listening on %s", addr)

	// Start the server and listen on the specified address
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
