package handler

import (
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteUserHandler(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	model := &model.User{}
	err := model.DeleteUser(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{
		"message": "User has been deleted",
	})
}
