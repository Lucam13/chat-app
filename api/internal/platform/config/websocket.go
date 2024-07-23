package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type (
	// WebSocketConfig is a configuration for web socket connections.
	WebSocketConfig struct {
		Upgrader websocket.Upgrader
	}
)

func NewWebSocketConfig(ctx *gin.Context) *WebSocketConfig {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &WebSocketConfig{
		Upgrader: upgrader,
	}
}

// GetUpgrader returns the web socket upgrader.
func (c WebSocketConfig) GetUpgrader() *websocket.Upgrader {
	return &c.Upgrader
}
