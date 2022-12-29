package controllers

import (
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func LibraryController(router *gin.Engine) {
	router.GET("/GetAllBooks",routes.GetAllBooks)
	router.POST("/AddBook",routes.AddBook)
	router.POST("/DeleteBook/:id",routes.DeleteBook)
}
