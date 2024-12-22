package entity

import (
	constants "NotesBuddy/Constants"
	database "NotesBuddy/Database"
	dto "NotesBuddy/Model/Dto"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

type User struct {
	Name      string
	Email     string
	Phone     string
	TechStack []dto.TechStack
	College   string
	Age       int
}

func (request *User) CreateUser(ctx *context.Context) error {
	collection, err := database.ConnectDB(ctx, constants.COLLECTION_USERS)
	if err != nil {
		return errors.New(err.Error())
	}

	document := bson.D{
		{Key: "name", Value: request.Name},
		{Key: "email", Value: request.Email},
		{Key: "phone", Value: request.Phone},
		{Key: "tech_stack", Value: request.TechStack},
		{Key: "college", Value: request.College},
		{Key: "age", Value: request.Age},
	}

	_, er := collection.InsertOne(*ctx, document)
	if er != nil {
		return er
	}

	return nil
}

func GetUserDetails(ctx *context.Context, phone string) (*dto.User, error) {
	collection, err := database.ConnectDB(ctx, constants.COLLECTION_USERS)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"phone": phone,
	}

	res := collection.FindOne(*ctx, filter)

	var user dto.User
	er := res.Decode(&user)
	if er != nil {
		return nil, er
	}

	return &user, nil
}
