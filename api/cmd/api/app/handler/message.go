package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lean1097/chat-back/cmd/api/app/handler/command"
	messageService "github.com/lean1097/chat-back/internal/chatapp/message/service"
)

type (
	MessageHandler struct {
		messageService messageService.MessageService
	}
)

func NewMessageHandler(messageService messageService.MessageService) MessageHandler {
	return MessageHandler{
		messageService: messageService,
	}
}

func (h MessageHandler) GetByChatID() gin.HandlerFunc {
	return func(c *gin.Context) {
		chatIDStr := c.Param("chat_id")
		chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chat ID"})
			return
		}

		messages, err := h.messageService.GetByChatID(c.Request.Context(), chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, messages)
	}
}

func (h MessageHandler) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req command.MessageCommand
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := h.messageService.Save(c.Request.Context(), req.Text, req.UserID, req.ChatID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusCreated)
	}
}
