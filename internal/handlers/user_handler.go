package handlers

import (
	"IDS/api/internal/services"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		service: services.NewUserService(db),
	}
}

func (h *UserHandler) GetUserRequests(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	requests, err := h.service.GetUserRequests(id)
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get requests"})
		return
	}

	if requests == nil {
		c.JSON(http.StatusOK, []any{})
		return
	}

	c.JSON(http.StatusOK, requests)
}
