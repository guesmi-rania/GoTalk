package main

import (
	"fmt"
	"net/http"
	"os"

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
	// Connect to DB
	database.Connect()

	// Migrate tables
	user.Migrate(database.DB)
	message.Migrate(database.DB)

	// Init hub
	hub := ws.NewHub()
	go hub.Run()

	// Init Gin
	r := gin.Default()

	// Health check
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

	// 🔹 PORT pour Render
	port := os.Getenv("PORT") // Render fournit ce port
	if port == "" {
		port = "8080" // fallback pour dev local
	}

	fmt.Println("Serving on port", port)
	r.Run("0.0.0.0:" + port) // écoute sur toutes les interfaces
}
