package base

import (
	database "Hackmate/Database"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var DatabaseInstance *mongo.Client

func Initiate() {
	ctx := context.Background()
	databaseInstance, err := database.ConnectDB(&ctx)
	if err != nil {
		fmt.Println("db connection error")
	}
	DatabaseInstance = databaseInstance
}
