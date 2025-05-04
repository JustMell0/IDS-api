package handlers

import (
	"IDS/api/internal/models"
	"IDS/api/internal/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	service *services.BookingService
}

func NewBookingHandler(db *sql.DB) *BookingHandler {
	return &BookingHandler{
		service: services.NewBookingService(db),
	}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	booking := models.Booking{}
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking data"})
		return
	}

	guest, err := h.service.CreateBooking(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, guest)
}
