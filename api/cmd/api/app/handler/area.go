package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lean1097/chat-back/cmd/api/app/handler/command"
	areaService "github.com/lean1097/chat-back/internal/chatapp/area/service"
)

type (
	// AreaHandler is a handler for area related operations.
	AreaHandler struct {
		areaService areaService.AreaService
	}
)

func NewAreaHandler(areaService areaService.AreaService) AreaHandler {
	return AreaHandler{
		areaService: areaService,
	}
}

func (h *AreaHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		areas, err := h.areaService.Get(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, areas)
	}
}

func (h *AreaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		areaIDStr := c.Param("id")
		areaID, err := strconv.ParseInt(areaIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid area ID"})
			return
		}

		area, err := h.areaService.GetByID(c.Request.Context(), areaID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, area)
	}
}

func (h *AreaHandler) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req command.AreaCommand
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := h.areaService.Save(c.Request.Context(), req.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusCreated)
	}
}

func (h *AreaHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		areaIDStr := c.Param("id")
		areaID, err := strconv.ParseInt(areaIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid area ID"})
			return
		}

		if err := h.areaService.Delete(c.Request.Context(), areaID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
