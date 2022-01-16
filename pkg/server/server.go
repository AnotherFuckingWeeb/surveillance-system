package server

import (
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/handler"
	CameraController "github.com/AnotherFuckingWeeb/surveillance-system/pkg/handler/camera"
	UserController "github.com/AnotherFuckingWeeb/surveillance-system/pkg/handler/user"
	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	router := gin.New()

	router.Use(middleware.CORS())

	//router.Static("public/assets/fonts", "../public/assets/fonts")

	router.POST("/login/", handler.LoginHandler)
	router.POST("/signup/", handler.SignUpHandler)

	authorized := router.Group("/api")

	authorized.GET("/me/", middleware.AuthorizeJWT(), handler.Me)
	authorized.GET("/camera/:id/", middleware.AuthorizeJWT(), CameraController.GetCameraHandler)

	protected := router.Group("/api/admin")

	protected.GET("/users/", middleware.AuthorizeJWT(), middleware.AuthorizeAdmin(), UserController.GetUsersHandler)
	protected.GET("/user/:id/", middleware.AuthorizeJWT(), middleware.AuthorizeAdmin(), UserController.GetUserHandler)
	protected.POST("/user/", middleware.AuthorizeJWT(), middleware.AuthorizeAdmin(), UserController.AddUserHandler)
	protected.PUT("/user/:id/", middleware.AuthorizeJWT(), middleware.AuthorizeAdmin(), UserController.UpdateUserHandler)
	protected.DELETE("/user/:id/", middleware.AuthorizeJWT(), middleware.AuthorizeAdmin(), UserController.DeleteUserHandler)

	protected.GET("/cameras/", middleware.AuthorizeJWT(), middleware.AuthorizeAdmin(), CameraController.GetCamerasHandler)
	protected.GET("/camera/:id/", middleware.AuthorizeJWT(), middleware.AuthorizeAdmin(), CameraController.GetCameraHandler)
	protected.POST("/camera/", middleware.AuthorizeJWT(), middleware.AuthorizeAdmin(), CameraController.AddCameraHandler)
	protected.PUT("/camera/:id/", middleware.AuthorizeJWT(), middleware.AuthorizeAdmin(), CameraController.UpdateCameraHandler)
	protected.DELETE("/camera/:id/", middleware.AuthorizeJWT(), middleware.AuthorizeAdmin(), CameraController.DeleteCameraHandler)

	return router
}
