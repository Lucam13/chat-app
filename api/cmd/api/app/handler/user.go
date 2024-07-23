package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lean1097/chat-back/internal/chatapp/user"
	userService "github.com/lean1097/chat-back/internal/chatapp/user/service"
)

type (
	// UserHandler is a handler for user-related requests.
	UserHandler struct {
		userService userService.UserService
	}
)

// NewUserHandler creates a new UserHandler with the given dependencies.
func NewUserHandler(userService userService.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

// Get gets all users.
func (h UserHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := h.userService.Get(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

// GetByID gets a user by ID.
func (h UserHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.Param("id")
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
			return
		}

		user, err := h.userService.GetByID(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// Save saves a new user.
func (h UserHandler) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user user.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := h.userService.Save(c.Request.Context(), user.Username, user.Rol, user.AreaID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusCreated)
	}
}

// Delete deletes a user by ID.
func (h UserHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.Param("id")
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
			return
		}

		if err := h.userService.Delete(c.Request.Context(), userID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
