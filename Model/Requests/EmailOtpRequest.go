package requests

import (
	utils "Hackmate/Utils"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type EmailOtpRequest struct {
	Email string `json:"email"`
	Otp   int    `json:"otp"`
}

func (request *EmailOtpRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.BindJSON(request); err != nil {
		return &ctx, err
	}
	return &ctx, nil
}

func (request *EmailOtpRequest) Validate(ctx *context.Context) error {
	if !utils.IsValidEmail(request.Email) {
		return errors.New("enter a valid mail")
	}

	if request.Otp >= 1000000 || request.Otp <= 100000 {
		return errors.New("invalid otp")
	}
	return nil
}
