package response

import (
	hmerrors "Hackmate/Model/Errors"
	"context"
)

type StringResponse struct {
	Url string `json:"url"`
	BaseResponse
}

func (response *StringResponse) Fail(ctx *context.Context, key string, error *hmerrors.Bderror, request interface{}) *StringResponse {
	response.Status = false
	response.Message = key
	response.Error = error
	response.Request = request
	return response
}

func (response *StringResponse) Success(ctx *context.Context, key string, url string) *StringResponse {
	response.Status = true
	response.Message = key
	response.Url = url
	return response
}
