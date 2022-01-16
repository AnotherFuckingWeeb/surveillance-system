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

func SignUpHandler(c *gin.Context) {
	var (
		newUser model.User
		err     error
	)

	err = c.BindJSON(&newUser)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	cryptPassword, _ := hashing.HashPassword(newUser.Password)

	userModel := model.User{
		DNI:      newUser.DNI,
		Role:     1,
		Name:     newUser.Name,
		Lastname: newUser.Lastname,
		Password: cryptPassword,
	}

	user, err := userModel.Create()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
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
