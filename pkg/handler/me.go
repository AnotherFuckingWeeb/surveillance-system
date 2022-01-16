package handler

import (
	"net/http"

	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/auth"
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	token, err := auth.ValidateJWT(c)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})

		return
	}

	user := &model.User{}

	if token.Valid {
		claims := token.Claims.(*auth.UserClaims)

		user.ID = claims.ID
		user.Role = claims.Role
		user.Name = claims.Name
		user.Lastname = claims.Lastname
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "you're not logeed in",
		})

		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"user": user,
	})
}
