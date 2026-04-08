package handler

import (
	"net/http"
	models "tipes/internal/model"
	"tipes/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return AuthHandler{service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user models.User

	user.FirstName = c.PostForm("firstname")
	user.LastName = c.PostForm("lastname")
	user.Email = c.PostForm("email")
	user.NoTelephone = c.PostForm("no_telephone")
	user.Password = c.PostForm("password")

	err := h.service.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register berhasil",
	})
}
