package main

import (
	"go-api/controllers"
	databaseconfiguration "go-api/databaseConfiguration"
	"go-api/environmentVariables"

	"github.com/gin-gonic/gin"
)

func init() {
	environmentVariables.LoadEnvironmentVariables()
	databaseconfiguration.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	controllers.UserController(router)
	controllers.LibraryController(router)
	router.Run()
}
