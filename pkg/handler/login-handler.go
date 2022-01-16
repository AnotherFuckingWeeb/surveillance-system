package handler

import (
	"net/http"
	"time"

	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/auth"
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/hashing"
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginBody struct {
	DNI      int
	Password string
}

func LoginHandler(c *gin.Context) {
	var (
		userModel model.User
		body      LoginBody
		err       error
	)

	err = c.BindJSON(&body)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	user, err := userModel.GetUserByDNI(body.DNI)

	if user == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})

		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	if !hashing.CheckPassword(body.Password, user.Password) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"message": "Password does not match",
		})

		return
	}

	token := auth.GenerateJWT(&auth.UserClaims{
		ID:       user.ID,
		Role:     user.Role,
		Name:     user.Name,
		Lastname: user.Lastname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "surveillance-system",
		},
	})

	c.IndentedJSON(http.StatusOK, gin.H{
		"token": token,
	})
}
