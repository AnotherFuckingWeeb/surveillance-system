package handler

import (
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteCameraHandler(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	model := &model.Camera{}

	err := model.DeleteCamera(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Camera Has Been Deleted",
	})
}
