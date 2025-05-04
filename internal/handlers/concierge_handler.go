package handlers

import (
	"IDS/api/internal/services"
	"database/sql"
	"log"
	"net/http"
	"strconv"

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

func (h *ConciergeHandler) AcceptRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	err = h.service.AcceptRequest(id)
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to accept request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request successfully accepted"})
}

func (h *ConciergeHandler) RejectRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	err = h.service.RejectRequest(id)
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reject request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request successfully rejected"})
}

func (h *ConciergeHandler) GetInHotelServices(c *gin.Context) {
	services, err := h.service.GetInHotelServices()
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get in-hotel services"})
		return
	}

	c.JSON(http.StatusOK, services)
}

func (h *ConciergeHandler) GetTourServices(c *gin.Context) {
	services, err := h.service.GetTourServices()
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tour services"})
		return
	}

	c.JSON(http.StatusOK, services)
}
