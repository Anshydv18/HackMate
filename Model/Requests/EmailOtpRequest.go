package requests

import (
	hmerrors "Hackmate/Model/Errors"
	utils "Hackmate/Utils"
	"context"

	"github.com/gin-gonic/gin"
)

type EmailOtpRequest struct {
	Email string `json:"email"`
	Otp   int    `json:"otp"`
}

func (request *EmailOtpRequest) Initiate(c *gin.Context, key string) (*context.Context, *hmerrors.Bderror) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.BindJSON(request); err != nil {
		return &ctx, hmerrors.InvalidInputError(&ctx, key, request)
	}
	return &ctx, nil
}

func (request *EmailOtpRequest) Validate(ctx *context.Context) *hmerrors.Bderror {
	if !utils.IsValidEmail(request.Email) {
		return hmerrors.InvalidInputError(ctx, "enter a valid mail", request)
	}

	if request.Otp >= 1000000 || request.Otp <= 100000 {
		return hmerrors.InvalidInputError(ctx, "invalid otp", request)
	}
	return nil
}
