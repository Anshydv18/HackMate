package services

import (
	entity "NotesBuddy/Model/Entity"
	redisentity "NotesBuddy/Model/RedisEntity"
	requests "NotesBuddy/Model/Requests"
	"context"
	"errors"
)

func CreateUserProfile(ctx *context.Context, request *requests.UserProfileRequest) error {

	userDto, _ := redisentity.GetUserFromCache(ctx, request.Phone)
	if userDto != nil {
		return errors.New("user already exits")
	}
	AlreadyExits, err := entity.IsUserAlreadyExists(ctx, request.Phone)
	if err != nil {
		return err
	}

	if AlreadyExits {
		return errors.New("user already exits")
	}
	UserEntity := entity.User{
		Name:      request.Name,
		College:   request.College,
		Age:       request.Age,
		TechStack: request.TechStacks,
		Phone:     request.Phone,
		Email:     request.Email,
	}

	return UserEntity.CreateUser(ctx)
}
