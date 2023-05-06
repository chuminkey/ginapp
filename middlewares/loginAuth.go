package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	accounts := gin.Accounts{
		"admin": "admin123",
	}
	return gin.BasicAuth(accounts)
}
