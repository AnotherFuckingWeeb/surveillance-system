package middleware

import (
	"net/http"

	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/auth"
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
)

func AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := auth.ValidateJWT(c)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		user := model.User{}

		if token.Valid {
			claims := token.Claims.(*auth.UserClaims)
			user.Role = claims.Role
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if user.Role < 1 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
