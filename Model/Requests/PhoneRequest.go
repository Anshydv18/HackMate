package requests

import (
	hmerrors "Hackmate/Model/Errors"
	utils "Hackmate/Utils"
	"context"

	"github.com/gin-gonic/gin"
)

type PhoneRequest struct {
	Phone string `json:"phone"`
}

func (request *PhoneRequest) Initiate(c *gin.Context, key string) (*context.Context, *hmerrors.Bderror) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, hmerrors.InvalidInputError(&ctx, err.Error(), request)
	}

	return &ctx, nil
}

func (request *PhoneRequest) Validate(ctx *context.Context) *hmerrors.Bderror {
	if len(request.Phone) < 10 {
		return hmerrors.InvalidInputError(ctx, "number size cannot be less than 10", request)
	}

	if !utils.IsValidPhone(request.Phone) {
		return hmerrors.InvalidInputError(ctx, "not a valid phone", request)
	}
	return nil
}
