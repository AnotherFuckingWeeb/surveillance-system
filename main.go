package main

import (
	"os"

	"github.com/AnotherFuckingWeeb/surveillance-system/pkg/server"
)

func main() {
	port := os.Getenv("PORT")

	app := server.Server()
	app.Run(":" + port)
}
