package services

import (
	dto "Hackmate/Model/Dto"
	entity "Hackmate/Model/Entity"
	redisentity "Hackmate/Model/RedisEntity"
	requests "Hackmate/Model/Requests"
	utils "Hackmate/Utils"
	"context"
	"crypto/rand"
	"errors"
	"math/big"
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
		Name:       request.Name,
		College:    request.College,
		Age:        request.Age,
		TechStacks: request.TechStacks,
		Phone:      request.Phone,
		Email:      request.Email,
	}

	go redisentity.SetUserCache(ctx, request.Phone, &UserDto)

	return &UserDto, nil
}

func Login(ctx *context.Context, phone string) (*dto.User, error) {
	userData, _ := redisentity.GetUserFromCache(ctx, phone)
	if userData != nil {
		return userData, nil
	}

	data, err := entity.GetUserDetails(ctx, phone)
	if err != nil {
		return nil, err
	}

	go redisentity.SetUserCache(ctx, data.Phone, data)
	return data, nil
}

func GetUserByEmail(ctx *context.Context, email string) (*dto.User, error) {
	if !utils.IsValidEmail(email) {
		return nil, errors.New("not a valid email")
	}

	userData, _ := redisentity.GetUserFromCache(ctx, email)
	if userData != nil {
		return userData, nil
	}

	data, err := entity.GetUserDetailsWithEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	go redisentity.SetUserCache(ctx, data.Email, data)
	return data, nil
}

func VerifyUserOtp(ctx *context.Context, email string, otp int) (bool, error) {
	storedOtp, _ := redisentity.GetOtpCache(ctx, email)
	if storedOtp != 0 {
		return storedOtp == otp, nil
	}
	return false, errors.New("otp expired")
}

func GenerateUserOtp(ctx *context.Context, email string) error {
	Otp, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return err
	}
	otp := Otp.Int64() + 100000

	go SendOtpMail(ctx, email, otp)
	return redisentity.SetOtpCache(ctx, email, int(otp))
}
