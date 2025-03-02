package requests

import (
	dto "Hackmate/Model/Dto"
	Error "Hackmate/Model/Errors"
	"context"
	"strings"

	"github.com/gin-gonic/gin"
)

type HackPostRequest struct {
	*dto.HackathonPost
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
