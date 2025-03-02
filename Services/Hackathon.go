package services

import (
	base "Hackmate/Base"
	constants "Hackmate/Constants"
	dto "Hackmate/Model/Dto"
	hmerrors "Hackmate/Model/Errors"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateHackathonPost(ctx *context.Context, request *dto.HackathonPost) *hmerrors.Bderror {
	dbClient := base.DatabaseInstance
	if dbClient == nil {
		return hmerrors.DataBaseConnectionError(ctx, "connection error with db")
	}

	collection := dbClient.Database(constants.DB_NAME).Collection(constants.COLLECTION_USERS)
	currentTime := primitive.NewDateTimeFromTime(time.Now())

	request.Id = primitive.NewObjectID()
	request.CreatedAt = currentTime
	request.UpdatedAt = currentTime

	_, err := collection.InsertOne(*ctx, request)
	if err != nil {
		return hmerrors.InvalidInputError(ctx, err.Error())
	}

	return nil
}

func FetchAllTrendingPost(ctx *context.Context) ([]*dto.HackathonPost, *hmerrors.Bderror) {
	dbClient := base.DatabaseInstance
	if dbClient == nil {
		return nil, hmerrors.DataBaseConnectionError(ctx, "error in connection")
	}

	collection := dbClient.Database(constants.DB_NAME).Collection(constants.COLLECTION_POSTS)
	options := options.Find()
	options.SetSort(bson.M{"created_at": -1})
	options.SetLimit(10)

	cursor, err := collection.Find(*ctx, bson.M{}, options)
	if err != nil {
		return nil, hmerrors.DataBaseReadError(ctx, err.Error())
	}
	var posts []*dto.HackathonPost
	if err := cursor.All(*ctx, &posts); err != nil {
		return nil, hmerrors.DataBaseReadError(ctx, err.Error())
	}
	return posts, nil
}
