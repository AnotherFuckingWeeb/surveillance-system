package handler

import (
	"net/http"
	"strconv"

	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
)

func GetCameraHandler(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

	model := &model.Camera{}
	camera, err := model.GetCameraById(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"camera": camera,
	})
}
