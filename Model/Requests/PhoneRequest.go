package requests

import (
	utils "NotesBuddy/Utils"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type PhoneRequest struct {
	Phone string `json:"phone"`
}

func (request *PhoneRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.BindJSON(&request); err != nil {
		return nil, errors.New("failed while binding request")
	}

	return &ctx, nil
}

func (request *PhoneRequest) Validate(ctx *context.Context) error {
	if len(request.Phone) < 10 {
		return errors.New(" number size cannot be less than 10")
	}

	if !utils.IsValidPhone(request.Phone) {
		return errors.New("enter a valid number")
	}
	return nil
}
