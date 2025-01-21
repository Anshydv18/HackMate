package requests

import (
	hmerrors "Hackmate/Model/Errors"
	"context"

	"github.com/gin-gonic/gin"
)

type MailRequest struct {
	Mail           []string `json:"mail"`
	SenderName     string   `json:"sender_name"`
	ContactDetails string   `json:"contact_details"`
	TeamName       string   `json:"team_name"`
	Status         int      `json:"status"`
}

func (request *MailRequest) Initiate(c *gin.Context, key string) (*context.Context, *hmerrors.Bderror) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, hmerrors.InvalidInputError(&ctx, key, request)
	}

	return &ctx, nil
}

func (request *MailRequest) Validate(ctx *context.Context) *hmerrors.Bderror {

	if len(request.Mail) == 0 {
		return hmerrors.InvalidInputError(ctx, "reciever mail cannot be empty", request)
	}
	return nil
}
