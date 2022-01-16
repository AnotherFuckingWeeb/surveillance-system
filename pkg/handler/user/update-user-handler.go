package handler

import (
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateUserHandler(c *gin.Context) {
	var (
		updatedUser model.User
		err         error
	)

	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

	err = c.BindJSON(&updatedUser)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	userModel := &model.User{
		DNI:      updatedUser.DNI,
		Role:     updatedUser.Role,
		Name:     updatedUser.Name,
		Lastname: updatedUser.Lastname,
		Password: updatedUser.Password,
	}

	err = userModel.UpdateUser(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{
		"message": "User Has Been Updated",
	})
}
