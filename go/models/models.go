package models

type Server struct {
	ClientInfo map[string]string // Client IP addresses and corresponding id
	ClientIps  []string
}

// Used to Register the client to connect to
// the websocket/signaling server
type Client struct {
	Request  string `json:"request"` // The client request :- Join or Create
	Username string `json:"username"`
}
