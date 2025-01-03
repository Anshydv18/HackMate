package requests

import (
	dto "Hackmate/Model/Dto"
	utils "Hackmate/Utils"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type UserProfileRequest struct {
	Name          string          `json:"name"`
	College       string          `json:"college"`
	TechStacks    []dto.TechStack `json:"tech_stacks"`
	Phone         string          `json:"phone_number"`
	Email         string          `json:"email"`
	Age           int             `json:"age"`
	GithubLink    string          `json:"github_link"`
	PortfolioLink string          `json:"portfolio_link"`
	ProfilePhoto  string          `json:"profile_photo"`
}

func (request *UserProfileRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	return &ctx, nil
}

func (request *UserProfileRequest) Validate(ctx *context.Context) error {
	if len(request.Name) < 3 {
		return errors.New("enter a valid name")
	}

	if len(request.College) <= 3 {
		return errors.New("enter college name")
	}

	if len(request.TechStacks) <= 0 {
		return errors.New("enter tech stack")
	}

	if !utils.IsValidPhone(request.Phone) {
		return errors.New("enter a valid phone number")
	}

	if !utils.IsValidEmail(request.Email) {
		return errors.New("enter a valid email")
	}

	if request.Age <= 5 {
		return errors.New("age should be greater than 5")
	}

	return nil
}
