package routes

import (
	"github.com/gin-gonic/gin"
)

func MainRoute(r *gin.Engine) {
	AuthRoute(r)
}
