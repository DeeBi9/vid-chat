package client

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/Deepanshuisjod/vid-chat/models"
	"github.com/gorilla/websocket"
)

type SliceID struct {
	mu sync.Mutex
	id []string
}

var (
	addr = "ws://localhost:7071/ws"
)

func GenID() string {
	var UniqueId string = ""
	var IDlength int = 6

	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < IDlength; i++ {
		char := charset[rand.Intn(len(charset))]
		UniqueId = UniqueId + string(char)
	}

	fmt.Println(string(UniqueId))
	return UniqueId
}

func (s *SliceID) Append(roomid string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.id = append(s.id, roomid)
}

func (s *SliceID) Remove(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	index := 0
	for {
		if id == s.id[index] {
			s.id = append(s.id[:index], s.id[index+1:]...)
			break
		}
		index += 1
	}
}

// Connect is used to connect to the Socket Server
// Connect act as the Socket Client
func Connect(rw http.ResponseWriter, r *http.Request, client *models.Client) {
	dialer, _, err := websocket.DefaultDialer.Dial(addr, r.Header)
	if err != nil {
		http.Error(rw, err.Error()+"[Error]", http.StatusInternalServerError)
		return
	}
	defer dialer.Close()

	done := make(chan struct{})

	slice := &SliceID{
		id: []string{},
	}

	go func() {
		defer close(done)

		for {
			_, message, err := dialer.ReadMessage()
			if err != nil {
				log.Println("error reading message: ", err)
				return
			}
			log.Printf("Recieved: %s", message)

			uniqueid := GenID()
			slice.Append(uniqueid)
			rw.Write([]byte(fmt.Sprintf("Your ROOM ID is :%s", uniqueid)))

			go func(id string) {
				time.Sleep(time.Minute * 5)
				slice.Remove(id)
				rw.Write([]byte("Session will be closed after 5 minutes"))
			}(uniqueid)
		}
	}()

}
