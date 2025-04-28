package handlers

import (
	"IDS/api/internal/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReservationHandler struct {
	service *services.ReservationService
}

func NewReservationHandler(db *sql.DB) *ReservationHandler {
	return &ReservationHandler{
		service: services.NewReservationService(db),
	}
}

func (h *ReservationHandler) GetReservations(c *gin.Context) {
	reservations, err := h.service.GetReservations()
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reservations"})
		return
	}

	c.JSON(http.StatusOK, reservations)
}
