package services

import (
	dto "Hackmate/Model/Dto"
	entity "Hackmate/Model/Entity"
	hmerrors "Hackmate/Model/Errors"
	redisentity "Hackmate/Model/RedisEntity"
	requests "Hackmate/Model/Requests"
	utils "Hackmate/Utils"
	"context"
	"crypto/rand"
	"math/big"
)

func CreateUserProfile(ctx *context.Context, request *requests.UserProfileRequest) (*dto.User, *hmerrors.Bderror) {

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

func Login(ctx *context.Context, phone string) (*dto.User, *hmerrors.Bderror) {
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

func GetUserByEmail(ctx *context.Context, email string) (*dto.User, *hmerrors.Bderror) {
	if !utils.IsValidEmail(email) {
		return nil, hmerrors.InvalidInputError(ctx, "not a valid email", email)
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

func VerifyUserOtp(ctx *context.Context, email string, otp int) (bool, *hmerrors.Bderror) {
	storedOtp, _ := redisentity.GetOtpCache(ctx, email)
	if storedOtp != 0 {
		return storedOtp == otp, nil
	}
	return false, hmerrors.InvalidInputError(ctx, "otp expired", email)
}

func GenerateUserOtp(ctx *context.Context, email string) *hmerrors.Bderror {
	Otp, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return hmerrors.InvalidInputError(ctx, "otp", email)
	}
	otp := Otp.Int64() + 100000

	go SendOtpMail(ctx, email, otp)
	return redisentity.SetOtpCache(ctx, email, int(otp))
}
