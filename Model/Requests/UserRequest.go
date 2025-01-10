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
	TechStacks    []dto.TechStack `json:"tech_stacks,omitempty"`
	Phone         string          `json:"phone_number"`
	Email         string          `json:"email"`
	Age           int             `json:"age,omitempty"`
	GithubLink    string          `json:"github_link,omitempty"`
	PortfolioLink string          `json:"portfolio_link,omitempty"`
	ProfilePhoto  string          `json:"profile_photo,omitempty"`
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

	if !utils.IsValidPhone(request.Phone) {
		return errors.New("enter a valid phone number")
	}

	if !utils.IsValidEmail(request.Email) {
		return errors.New("enter a valid email")
	}

	if request.Age != 0 && request.Age <= 5 {
		return errors.New("age should be greater than 5")
	}

	return nil
}
