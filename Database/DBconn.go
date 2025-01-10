package database

import (
	constants "Hackmate/Constants"
	env "Hackmate/Env"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx *context.Context) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	UriEnv := env.Get(constants.MONGODBURI)
	option := options.Client().ApplyURI(UriEnv).SetServerAPIOptions(serverAPI)

	clients, err := mongo.Connect(*ctx, option)
	if err != nil {
		return nil, errors.New("mongodb connection failed")
	}

	return clients, nil
}
