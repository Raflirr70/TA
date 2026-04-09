package routes

import (
	"tipes/internal/config"
	"tipes/internal/handler"
	"tipes/internal/middleware"
	"tipes/internal/repository"
	"tipes/internal/service"

	"github.com/gin-gonic/gin"
)

func EmployeeRoute(r *gin.Engine) {
	empRepo := repository.NewEmployeeRepository(config.DB)
	empService := service.NewEmployeeService(empRepo)
	empHandler := handler.NewEmployeeHandler(empService)

	protected := r.Group("/employee")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/", empHandler.GetEmployees)            // ?branch_id=1
		protected.POST("/hire", empHandler.HireEmployee)       // hire
		protected.PUT("/edit", empHandler.EditEmployee)        // edit
		protected.DELETE("/fire/:id", empHandler.FireEmployee) // fire
	}
}
