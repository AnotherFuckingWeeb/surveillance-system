package handler

import (
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserHandler(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	userModel := &model.User{}
	user, err := userModel.GetUserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"user": user,
	})
}
