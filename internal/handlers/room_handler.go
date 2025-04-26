package handlers

import (
	"IDS/api/internal/services"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	service *services.RoomService
}

func NewRoomHandler(db *sql.DB) *RoomHandler {
	return &RoomHandler{
		service: services.NewRoomService(db),
	}
}

func (h *RoomHandler) GetRooms(c *gin.Context) {
	rooms, err := h.service.GetRooms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rooms"})
		return
	}

	c.JSON(http.StatusOK, rooms)
}
