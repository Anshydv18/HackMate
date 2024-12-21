package services

import (
	entity "NotesBuddy/Model/Entity"
	requests "NotesBuddy/Model/Requests"
	"context"
)

func CreateUserProfile(ctx *context.Context, request *requests.UserProfileRequest) error {
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
