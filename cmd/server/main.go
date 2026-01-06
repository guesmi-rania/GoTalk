package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"gotalk/internal/user"
	"gotalk/internal/message"
	ws "gotalk/internal/websocket"
	"gotalk/pkg/database"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	// 1️⃣ Connect to DB
	database.Connect()

	// 2️⃣ Migrate tables
	user.Migrate(database.DB)
	message.Migrate(database.DB)

	// 3️⃣ Init hub
	hub := ws.NewHub()
	go hub.Run()

	// 4️⃣ Init Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// WebSocket endpoint
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Upgrade error:", err)
			return
		}

		client := &ws.Client{
			Conn: conn,
			Send: make(chan []byte),
		}
		hub.Register <- client

		// Handle client read messages
		go func() {
			defer func() {
				hub.Unregister <- client
				conn.Close()
			}()
			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					break
				}
				hub.Broadcast <- msg
			}
		}()
	})

	// Start server
	r.Run(":8080")
}
