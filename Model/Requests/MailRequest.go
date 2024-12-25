package requests

import (
	utils "Hackmate/Utils"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

type MailRequest struct {
	Mail string `json:"mail"`
}

func (request *MailRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, errors.New("failed while binding request")
	}

	return &ctx, nil
}

func (request *MailRequest) Validate(ctx *context.Context) error {

	if !utils.IsValidEmail(request.Mail) {
		return errors.New("enter a valid email")
	}
	return nil
}
