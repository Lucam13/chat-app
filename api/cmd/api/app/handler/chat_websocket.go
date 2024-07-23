package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lean1097/chat-back/internal/chatapp/message"
	messageService "github.com/lean1097/chat-back/internal/chatapp/message/service"
	"github.com/lean1097/chat-back/internal/platform/config"
)

type (
	// ChatWebSocketHandler is a handler for chat web socket operations.
	ChatWebSocketHandler struct {
		messageService  messageService.MessageService
		websocketConfig *config.WebSocketConfig
	}
)

func NewChatWebSocketHandler(ms messageService.MessageService) ChatWebSocketHandler {
	return ChatWebSocketHandler{
		messageService: ms,
	}
}

func (h *ChatWebSocketHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new web socket configuration.
		h.websocketConfig = config.NewWebSocketConfig(c)

		// Upgrade the HTTP connection to a web socket connection.
		conn, err := h.websocketConfig.GetUpgrader().Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Error upgrading to web socket connection:", err)
			c.Error(err)
			return
		}
		defer conn.Close()

		// Read messages from the web socket connection.
		for {
			var msg message.Message
			err := conn.ReadJSON(&msg)
			if err != nil {
				fmt.Println("Error reading message from web socket connection:", err)
				c.Error(err)
				return
			}

			// Process the message.
			err = h.messageService.Save(c.Request.Context(), msg.Text, msg.UserID, msg.ChatID)
			if err != nil {
				fmt.Println("Error saving message:", err)
				c.Error(err)
				return
			}

			// Send the message to all connected clients.
			err = conn.WriteJSON(msg)
			if err != nil {
				fmt.Println("Error writing message to web socket connection:", err)
				c.Error(err)
				return
			}
		}
	}
}
