package handlers

import (
	"IDS/api/internal/services"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomTypeHandler struct {
	service *services.RoomTypeService
}

func NewRoomTypeHandler(db *sql.DB) *RoomTypeHandler {
	return &RoomTypeHandler{
		service: services.NewRoomTypeService(db),
	}
}

func (h *RoomTypeHandler) GetRoomTypes(c *gin.Context) {
	roomTypes, err := h.service.GetRoomTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rooms"})
		return
	}

	c.JSON(http.StatusOK, roomTypes)
}
