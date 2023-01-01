package routes

import (
	"fmt"
	databaseconfiguration "go-api/databaseConfiguration"
	"go-api/models"
	"go-api/redisCache"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {
	users := []models.User{}
	var url = ctx.Request.URL.Path
	users,err := redisCache.GetAll(url)
	if err==nil{
		ctx.JSON(200,users)
		return
	}
	result := databaseconfiguration.DB.Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	redisCache.SetAll(url, users)
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
	var name = ctx.Query("name")
	fmt.Println(name)
	url := ctx.Request.URL.Path + name
	fmt.Println(url)
	users,err := redisCache.Get(url)
	if err == nil {
		// got something as an answer
		ctx.JSON(200, users)
		return
	}
	result := databaseconfiguration.DB.Where("name = ?", name).Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	redisCache.Set(url, users)
	ctx.JSON(200, users)
}
