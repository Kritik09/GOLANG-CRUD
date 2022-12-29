package routes

import (
	databaseconfiguration "go-api/databaseConfiguration"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context){
	libraryBooks:=[]models.Library{}
	result:=databaseconfiguration.DB.Find(&libraryBooks)
	if result.Error!=nil{
		panic(result.Error)
	}
	ctx.JSON(200,&libraryBooks)
}
func AddBook(ctx *gin.Context){
	book:=models.Library{}
	err:=ctx.BindJSON(&book)
	if err!=nil{
		panic(err)
	}
	result:=databaseconfiguration.DB.Create(&book)
	if result.Error!=nil{
		panic(result.Error)
	}
	ctx.JSON(200,"Book Successfully Added to the Database")
}
func DeleteBook(ctx *gin.Context){
	var id=ctx.Param("id")
	result:=databaseconfiguration.DB.Delete(&models.Library{},id)
	if result.Error!=nil{
		panic(result.Error)
	}
	ctx.JSON(200,"Book Deleted Succesfully")
}