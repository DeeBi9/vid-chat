package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Deepanshuisjod/vid-chat/client"
	"github.com/Deepanshuisjod/vid-chat/models"
)

func JoinRoom(rw http.ResponseWriter, r *http.Request) {

}

func CreateRoom(rw http.ResponseWriter, r *http.Request) {
	// Register the client
	var registeredclient models.Client
	err := json.NewDecoder(r.Body).Decode(&registeredclient)
	if err != nil {
		http.Error(rw, "Error decoding Client information [JSON]", http.StatusInternalServerError)
		return
	}

	client.Connect(rw, r, &registeredclient)
}
