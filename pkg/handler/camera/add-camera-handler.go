package handler

import (
	"net/http"

	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
)

func AddCameraHandler(c *gin.Context) {
	var (
		newCamera model.Camera
		err       error
	)

	err = c.BindJSON(&newCamera)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	model := &model.Camera{
		Brand:       newCamera.Brand,
		CreatedAt:   newCamera.CreatedAt,
		Area:        newCamera.Area,
		Description: newCamera.Description,
	}

	modelErr := model.Create()

	if modelErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Camera has been created",
	})
}
