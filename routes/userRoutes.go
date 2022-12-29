package routes

import (
	databaseconfiguration "go-api/databaseConfiguration"
	"go-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {
	users := []models.User{}
	result := databaseconfiguration.DB.Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(200, &users)
}
func AddUser(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	result := databaseconfiguration.DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(200, "User Added Successfully")
}
func DeleteUser(ctx *gin.Context) {
	var id = ctx.Param("id")
	result := databaseconfiguration.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(200, "User Deleted Successfully")
}
func FindUserByName(ctx *gin.Context) {
	var name = ctx.Param("name")
	users := []models.User{}
	result := databaseconfiguration.DB.Where("name <> ?", name).Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(200, users)
}
