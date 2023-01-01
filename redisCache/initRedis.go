package redisCache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api/models"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func InitializeRedis() {
	ct := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Test@123",
		DB:       0,
	})
	client = ct
}

func Set(name string, userModel *models.User) {
	fmt.Println("Inside Set")
	json, err := json.Marshal(userModel)
	if err != nil {
		panic("Unable to marshall the userModel")
	}
	client.Set(context.Background(), name, json, 10*time.Second)
}

func Get(url string) (*models.User, error) {
	fmt.Println("Inside Get")
	result, err := client.Get(context.Background(), url).Result()
	user := models.User{}
	if err == redis.Nil {
		// There is no key of the given name so returning empty
		return &user, err
	}
	jsonErr := json.Unmarshal([]byte(result), &user)
	// err = json.Unmarshal([]byte(result), &result)
	if jsonErr != nil {
		panic("Unable to unmarshall")
	}
	return &user, nil
}

func SetAll(name string, userModel []models.User) {
	fmt.Println("Inside SetAll")
	json, err := json.Marshal(userModel)
	if err != nil {
		panic("Unable to marshall the userModel")
	}
	client.Set(context.Background(), name, json, 10*time.Second)
}
func GetAll(key string) ([]models.User,error){
	fmt.Println("Inside GetAll")
	result, err := client.Get(context.Background(), key).Result()
	var user []models.User
	if err == redis.Nil {
		// There is no key of the given name so returning empty
		return user, err
	}
	jsonErr := json.Unmarshal([]byte(result), &user)
	// err = json.Unmarshal([]byte(result), &result)
	if jsonErr != nil {
		panic("Unable to unmarshall")
	}
	return user, nil
}
