package requests

import (
	Error "Hackmate/Model/Errors"
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type HackPostRequest struct {
	Name          string    `json:"name"`
	DateTime      time.Time `json:"dateTime"`
	Location      string    `json:"location"`
	Description   string    `json:"description"`
	Theme         string    `json:"theme"`
	TeamSizeLimit int       `json:"teamSizeLimit"`
}

func (request *HackPostRequest) Initiate(c *gin.Context, key string) (*context.Context, *Error.Bderror) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.BindJSON(&request); err != nil {
		return &ctx, Error.InvalidInputError(&ctx, key, err)
	}
	return &ctx, nil
}

func (request *HackPostRequest) Validate(ctx *context.Context) *Error.Bderror {
	request.Name = strings.TrimSpace(request.Name)
	if len(request.Name) == 0 {
		return Error.InvalidInputError(ctx, "enter a valid name", request)
	}

	request.Theme = strings.TrimSpace(request.Theme)
	if len(request.Theme) == 0 {
		return Error.InvalidInputError(ctx, "enter a valid name", request)
	}

	if request.TeamSizeLimit == 0 {
		return Error.InvalidInputError(ctx, "enter a valid name", request)
	}
	return nil
}
