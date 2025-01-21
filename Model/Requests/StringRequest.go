package requests

import (
	hmerrors "Hackmate/Model/Errors"
	"context"

	"github.com/gin-gonic/gin"
)

type StringRequest struct {
	Key string `json:"key"`
}

func (request *StringRequest) Initiate(c *gin.Context, key string) (*context.Context, *hmerrors.Bderror) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.BindJSON(&request); err != nil {
		return &ctx, hmerrors.InvalidInputError(&ctx, key, request)
	}

	return &ctx, nil
}

func (request *StringRequest) Validate(ctx *context.Context) *hmerrors.Bderror {
	if len(request.Key) == 0 {
		return hmerrors.InvalidInputError(ctx, "key length cannot be 0", request)
	}
	return nil
}
