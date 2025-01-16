package requests

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type StringRequest struct {
	Key string `json:"key"`
}

func (request *StringRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.BindJSON(&request); err != nil {
		return &ctx, err
	}

	return &ctx, nil
}

func (request *StringRequest) Validate(ctx *context.Context) error {
	if len(request.Key) == 0 {
		return errors.New("key is empty")
	}
	return nil
}
