package base

import (
	database "Hackmate/Database"
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

var DatabaseInstance *mongo.Client
var RedisInstance *redis.Client
var CloudinaryInstance *cloudinary.Cloudinary

func Initiate() {
	ctx := context.Background()

	databaseInstance, err := database.ConnectDB(&ctx)
	if err != nil {
		fmt.Println("db connection error")
	}
	DatabaseInstance = databaseInstance

	RedisInstance = database.StartRedisServer()

	cld, err := database.CloudinaryConnect(&ctx)
	if err != nil {
		fmt.Println("cloudinary connection error")
	}
	CloudinaryInstance = cld
}
