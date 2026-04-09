package routes

import (
	"tipes/internal/config"
	"tipes/internal/handler"
	"tipes/internal/middleware"
	"tipes/internal/repository"
	"tipes/internal/service"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.Engine) {
	authRepo := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	// route
	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			c.JSON(200, gin.H{
				"user_id": userID,
			})
		})
	}
}
