package handler

import (
	"net/http"

	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/model"
	"github.com/gin-gonic/gin"
)

func GetCamerasHandler(c *gin.Context) {
	model := &model.Camera{}
	cameras, err := model.GetCameras()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"cameras": cameras,
	})
}
