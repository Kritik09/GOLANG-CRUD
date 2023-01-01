package controllers

import (
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func UserController(router *gin.Engine) {
	router.GET("/GetAllUser",routes.GetAllUser)
	router.POST("/AddUser",routes.AddUser)
	router.POST("/DeleteUser",routes.DeleteUser)
	router.GET("/FindUserByName",routes.FindUserByName)
}