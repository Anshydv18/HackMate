package requests

import (
	dto "Hackmate/Model/Dto"
	hmerrors "Hackmate/Model/Errors"
	utils "Hackmate/Utils"
	"context"

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

func (request *UserProfileRequest) Initiate(c *gin.Context, key string) (*context.Context, *hmerrors.Bderror) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, hmerrors.InvalidInputError(&ctx, err.Error(), request)
	}

	return &ctx, nil
}

func (request *UserProfileRequest) Validate(ctx *context.Context) *hmerrors.Bderror {
	if len(request.Name) < 3 {
		return hmerrors.InvalidInputError(ctx, "enter a valid name", request)
	}

	if len(request.College) <= 3 {
		return hmerrors.InvalidInputError(ctx, "enter college name", request)
	}

	if !utils.IsValidPhone(request.Phone) {
		return hmerrors.InvalidInputError(ctx, "enter a valid phone number", request)
	}

	if !utils.IsValidEmail(request.Email) {
		return hmerrors.InvalidInputError(ctx, "enter a valid email", request)
	}

	if request.Age != 0 && request.Age <= 5 {
		return hmerrors.InvalidInputError(ctx, "age should be greater than 5", request)
	}

	return nil
}
