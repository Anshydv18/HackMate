package services

import (
	dto "NotesBuddy/Model/Dto"
	entity "NotesBuddy/Model/Entity"
	redisentity "NotesBuddy/Model/RedisEntity"
	requests "NotesBuddy/Model/Requests"
	"context"
)

func CreateUserProfile(ctx *context.Context, request *requests.UserProfileRequest) (*dto.User, error) {

	UserEntity := entity.User{
		Name:      request.Name,
		College:   request.College,
		Age:       request.Age,
		TechStack: request.TechStacks,
		Phone:     request.Phone,
		Email:     request.Email,
	}

	if err := UserEntity.CreateUser(ctx); err != nil {
		return nil, err

	}
	UserDto := dto.User{
		Name:      request.Name,
		College:   request.College,
		Age:       request.Age,
		TechStack: request.TechStacks,
		Phone:     request.Phone,
		Email:     request.Email,
	}

	go redisentity.SetUserCache(ctx, &UserDto)

	return &UserDto, nil
}

func Login(ctx *context.Context, phone string) (*dto.User, error) {
	userData, _ := redisentity.GetUserFromCache(ctx, phone)
	if userData != nil {
		return userData, nil
	}

	data, err := entity.GetUserDetails(ctx, phone)
	go redisentity.SetUserCache(ctx, data)
	return nil, err
}
