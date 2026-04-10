package routes

import (
	"tipes/internal/config"
	"tipes/internal/handler"
	"tipes/internal/middleware"
	"tipes/internal/repository"
	"tipes/internal/service"
	"tipes/pkg/utils"

	"github.com/gin-gonic/gin"
)

func EmployeeRoute(r *gin.Engine) {
	empRepo := repository.NewEmployeeRepository(config.DB)
	empService := service.NewEmployeeService(empRepo)
	empHandler := handler.NewEmployeeHandler(empService)

	protected := r.Group("/employee")
	protected.Use(middleware.AuthMiddleware())

	{
		// ✅ semua user login boleh lihat
		protected.GET("/", empHandler.GetEmployees)

		// 🔥 hanya OWNER & MANAGER
		protected.POST("/hire",
			middleware.AuthorizeRoles(utils.RoleOwner, utils.RoleManager),
			empHandler.HireEmployee,
		)

		protected.PUT("/edit",
			middleware.AuthorizeRoles(utils.RoleOwner, utils.RoleManager),
			empHandler.EditEmployee,
		)

		protected.DELETE("/fire/:id",
			middleware.AuthorizeRoles(utils.RoleOwner, utils.RoleManager),
			empHandler.FireEmployee,
		)
	}
}
