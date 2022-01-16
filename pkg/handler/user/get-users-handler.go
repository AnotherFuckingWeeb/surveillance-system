package handler

import (
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsersHandler(c *gin.Context) {
	model := &model.User{}
	users, err := model.GetUsers()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"users": &users,
	})
}
