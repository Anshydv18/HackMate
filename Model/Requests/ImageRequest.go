package requests

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type ImageRequest struct {
	Image multipart.FileHeader `json:"image" form:"image"`
}

func (request *ImageRequest) Initiate(c *gin.Context, key string) (*context.Context, error) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.Bind(request); err != nil {
		return nil, errors.New("error while binding")
	}

	return &ctx, nil
}

func (request *ImageRequest) Validate(ctx *context.Context) error {
	if request.Image.Size > 3<<20 {
		return errors.New("image size should be less than 3 mb")
	}
	return nil
}
