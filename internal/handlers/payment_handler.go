package handlers

import (
	"IDS/api/internal/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *services.PaymentService
}

func NewPaymentHandler(db *sql.DB) *PaymentHandler {
	return &PaymentHandler{
		service: services.NewPaymentService(db),
	}
}

func (h *PaymentHandler) GetUnconfirmedPayments(c *gin.Context) {
	payments, err := h.service.GetUnconfirmedPayments()
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get payments"})
		return
	}

	c.JSON(http.StatusOK, payments)

}
