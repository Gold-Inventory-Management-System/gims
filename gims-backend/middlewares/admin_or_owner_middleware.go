package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeAdminOrOwner() gin.HandlerFunc {
	return func(c *gin.Context){
		role, _ := c.Get("role")
		if role != "admin" && role != "owner" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "only admin or owner can do this",
			})
			return
		}
		c.Next()
	}
}