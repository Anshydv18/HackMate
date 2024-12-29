package response

import (
	"context"
)

type StringResponse struct {
	Url string `json:"url"`
	BaseResponse
}

func (response *StringResponse) Fail(ctx *context.Context, key string, error string, request interface{}) *StringResponse {
	response.Status = false
	response.Message = key
	response.ErrorType = error
	response.Request = request
	return response
}

func (response *StringResponse) Success(ctx *context.Context, key string, url string) *StringResponse {
	response.Status = true
	response.Message = key
	response.Url = url
	return response
}
