package requests

import (
	"context"
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

func (request *HackPostRequest) Initiate(c *gin.Context, key string) (ctx *context.Context, errors error) {
	return nil, nil
}

func (request *HackPostRequest) Validate(ctx *context.Context) error {
	return nil
}
