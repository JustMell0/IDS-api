package handlers

import (
	"IDS/api/internal/services"
	"database/sql"
	"log"
	"net/http"
	"strconv"

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

func (h *ReservationHandler) GetUserReservations(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user" + c.Param("id")})
	}
	reservations, err := h.service.GetUserReservations(idParam)
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user reservations"})
		return
	}
	c.JSON(http.StatusOK, reservations)
}
