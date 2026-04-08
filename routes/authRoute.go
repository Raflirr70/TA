package routes

import (
	"tipes/internal/config"
	"tipes/internal/handler"
	"tipes/internal/repository"
	"tipes/internal/service"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.Engine) {
	authRepo := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	// route
	r.POST("/register", authHandler.Register)

	r.Run(":8080")
}
