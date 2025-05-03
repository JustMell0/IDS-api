package handlers

import (
	"IDS/api/internal/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConciergeHandler struct {
	service *services.ConciergeService
}

func NewConciergeHandler(db *sql.DB) *ConciergeHandler {
	return &ConciergeHandler{
		service: services.NewConciergeService(db),
	}
}

func (h *ConciergeHandler) GetRequests(c *gin.Context) {
	requests, err := h.service.GetRequests()
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get requests"})
		return
	}

	c.JSON(http.StatusOK, requests)
}
