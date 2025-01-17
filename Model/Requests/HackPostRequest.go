package requests

import (
	"context"
	"errors"
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

func (request *HackPostRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.BindJSON(&request); err != nil {
		return &ctx, nil
	}

	return &ctx, nil
}

func (request *HackPostRequest) Validate(ctx *context.Context) error {
	if len(request.Name) == 0 {
		return errors.New("enter the hackathon name")
	}

	return nil
}
