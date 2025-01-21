package entity

import (
	base "Hackmate/Base"
	constants "Hackmate/Constants"
	dto "Hackmate/Model/Dto"
	hmerrors "Hackmate/Model/Errors"

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

func (request *User) CreateUser(ctx *context.Context) *hmerrors.Bderror {
	dbclient := base.DatabaseInstance
	if dbclient == nil {
		return hmerrors.InvalidInputError(ctx, "client error", request)
	}

	collection := dbclient.Database(constants.DB_NAME).Collection(constants.COLLECTION_USERS)
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
		return hmerrors.InvalidInputError(ctx, er.Error(), request)
	}

	return nil
}

func GetUserDetails(ctx *context.Context, phone string) (*dto.User, *hmerrors.Bderror) {
	dbclient := base.DatabaseInstance
	if dbclient == nil {
		return nil, hmerrors.InvalidInputError(ctx, "client error", phone)
	}

	collection := dbclient.Database(constants.DB_NAME).Collection(constants.COLLECTION_USERS)

	filter := bson.M{
		"phone": phone,
	}
	res := collection.FindOne(*ctx, filter)

	var user dto.User
	er := res.Decode(&user)
	if er != nil {
		return nil, hmerrors.InvalidInputError(ctx, er.Error(), phone)
	}

	return &user, nil
}

func GetUserDetailsWithEmail(ctx *context.Context, email string) (*dto.User, *hmerrors.Bderror) {
	dbclient := base.DatabaseInstance
	if dbclient == nil {
		return nil, hmerrors.InvalidInputError(ctx, "client error", email)
	}

	collection := dbclient.Database(constants.DB_NAME).Collection(constants.COLLECTION_USERS)

	filter := bson.M{
		"email": email,
	}
	res := collection.FindOne(*ctx, filter)

	var user dto.User
	er := res.Decode(&user)
	if er != nil {
		return nil, hmerrors.InvalidInputError(ctx, er.Error(), email)
	}

	return &user, nil
}
