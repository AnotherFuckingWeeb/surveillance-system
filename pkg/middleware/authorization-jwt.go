package middleware

import (
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := auth.ValidateJWT(c)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if token.Valid {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
