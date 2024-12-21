package entity

import (
	database "NotesBuddy/Database"
	dto "NotesBuddy/Model/Dto"

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
	conn, err := database.ConnectDB(ctx)
	if err != nil {
		return err
	}

	collection := conn.Database("hackMate").Collection("users")

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
