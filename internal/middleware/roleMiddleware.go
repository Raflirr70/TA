package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeRoles(roles ...uint) gin.HandlerFunc {
	return func(c *gin.Context) {

		roleInterface, exists := c.Get("role_id")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "role tidak ditemukan"})
			return
		}

		roleID := roleInterface.(uint)

		for _, role := range roles {
			if roleID == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "akses ditolak"})
	}
}
