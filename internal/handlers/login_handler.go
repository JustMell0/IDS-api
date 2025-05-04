package handlers

import (
	"IDS/api/internal/models"
	"IDS/api/internal/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	service *services.LoginService
}

func NewLoginHandler(db *sql.DB) *LoginHandler {
	return &LoginHandler{
		service: services.NewLoginService(db),
	}
}

func (h *LoginHandler) Login(c *gin.Context) {
	login := models.Login{}
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	guest, err := h.service.Login(login.PhoneNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		log.Println(err)
		return
	}

	c.JSON(http.StatusFound, guest)
}
