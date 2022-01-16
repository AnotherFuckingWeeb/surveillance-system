package handler

import (
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/hashing"
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUserHandler(c *gin.Context) {
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

	encryptedPassword, _ := hashing.HashPassword(newUser.Password)

	model := &model.User{
		DNI:      newUser.DNI,
		Role:     0,
		Name:     newUser.Name,
		Lastname: newUser.Lastname,
		Password: encryptedPassword,
	}

	_, err = model.Create()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{
		"message": "User Has Been Created",
	})
}
