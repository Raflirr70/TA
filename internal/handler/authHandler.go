package handler

import (
	"net/http"
	models "tipes/internal/model"
	"tipes/internal/service"
	"tipes/pkg/utils"

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

	user.RoleID = 1
	user.BranchID = 1
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

func (h *AuthHandler) Login(c *gin.Context) {
	identifier := c.PostForm("identifier")
	password := c.PostForm("password")

	// print("email : ", email, "dan password : ", password)

	if identifier == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"messege": "Masukan email/no telephone dan password"})
		return
	}

	//manggil print

	user, err := h.service.Login(identifier, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.UserID, user.Email, user.RoleID, user.BranchID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token":   token,
		// "user": gin.H{
		// 	"id":          user.UserID,
		// 	"name":        user.Name,
		// 	"email":       user.Email,
		// 	"noTelephone": user.NoTelephone,
		// },
	})

}
