package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lean1097/chat-back/cmd/api/app/handler/command"
	chatService "github.com/lean1097/chat-back/internal/chatapp/chat/service"
)

type (
	// ChatHandler is a handler for chat operations.
	ChatHandler struct {
		chatService chatService.ChatService
	}
)

func NewChatHandler(chatService chatService.ChatService) ChatHandler {
	return ChatHandler{
		chatService: chatService,
	}
}

func (h *ChatHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		chats, err := h.chatService.Get(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, chats)
	}
}

func (h *ChatHandler) GetByAreaID() gin.HandlerFunc {
	return func(c *gin.Context) {
		areaIDStr := c.Param("id")
		areaID, err := strconv.ParseInt(areaIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid area ID"})
			return
		}

		chats, err := h.chatService.GetByAreaID(c.Request.Context(), areaID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, chats)
	}
}

func (h *ChatHandler) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req command.ChatCommand
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.chatService.Save(c.Request.Context(), req.FromAreaID, req.ToAreaID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusCreated)
	}
}

func (h *ChatHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		chatIDStr := c.Param("id")
		chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chat ID"})
			return
		}

		err = h.chatService.Delete(c.Request.Context(), chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
