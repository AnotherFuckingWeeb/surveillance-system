package handler

import (
	"net/http"
	"strconv"

	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
)

func UpdateCameraHandler(c *gin.Context) {
	var (
		updatedCamera model.Camera
		err           error
	)

	id, _ := strconv.ParseInt(c.Param("id"), 0, 64)

	err = c.BindJSON(&updatedCamera)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	cameraModel := &model.Camera{
		Brand:       updatedCamera.Brand,
		Area:        updatedCamera.Area,
		Description: updatedCamera.Description,
	}

	modelErr := cameraModel.UpdateCamera(id)

	if modelErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Camera Has Been Updated",
	})
}
