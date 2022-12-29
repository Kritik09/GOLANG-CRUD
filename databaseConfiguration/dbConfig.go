package databaseconfiguration

import (
	"fmt"
	"go-api/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	db_host:=os.Getenv("DB_HOST")
	db_user:=os.Getenv("DB_USER")
	db_password:=os.Getenv("DB_PASSWORD")
	db_name:=os.Getenv("DB_NAME")
	db_port:=os.Getenv("DB_PORT")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",db_host,db_user,db_password,db_name,db_port,
	)
	// dsn := "host=localhost user=admin password=admin dbname=library port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Library{})
	DB = db
}
