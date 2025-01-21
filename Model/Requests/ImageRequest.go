package requests

import (
	hmerrors "Hackmate/Model/Errors"
	"context"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type ImageRequest struct {
	Image multipart.FileHeader `json:"image" form:"image"`
}

func (request *ImageRequest) Initiate(c *gin.Context, key string) (*context.Context, *hmerrors.Bderror) {
	_ctx, _ := c.Get("context")
	ctx := _ctx.(context.Context)

	if err := c.Bind(request); err != nil {
		return nil, hmerrors.InvalidInputError(&ctx, key, request)
	}

	return &ctx, nil
}

func (request *ImageRequest) Validate(ctx *context.Context) *hmerrors.Bderror {
	if request.Image.Size > 3<<20 {
		return hmerrors.InvalidInputError(ctx, "image size should be less than 3 mb", request)
	}
	return nil
}
