package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Deepanshuisjod/vid-chat/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", handlers.CreateRoom)
	postRouter.HandleFunc("/joincall", handlers.JoinRoom)

	server := &http.Server{
		Addr:         ":7070",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second, // Increased timeout
		WriteTimeout: 5 * time.Second, // Increased timeout
	}

	log.Fatal(server.ListenAndServe())

}
