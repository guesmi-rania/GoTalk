package websocket

import "github.com/gorilla/websocket"

type Client struct {
    Conn *websocket.Conn
    Send chan []byte
}

type Hub struct {
    Register   chan *Client
    Unregister chan *Client
    Broadcast  chan []byte
    Clients    map[*Client]bool
}

func NewHub() *Hub {
    return &Hub{
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Broadcast:  make(chan []byte),
        Clients:    make(map[*Client]bool),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.Register:
            h.Clients[client] = true
        case client := <-h.Unregister:
            delete(h.Clients, client)
        case message := <-h.Broadcast:
            for client := range h.Clients {
                client.Send <- message
            }
        }
    }
}
