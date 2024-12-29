package response

import (
	dto "Hackmate/Model/Dto"
	"context"
)

type UserResponse struct {
	Data *dto.User `json:"user"`
	BaseResponse
}

func (response *UserResponse) Fail(ctx *context.Context, Error error, key string, request interface{}) *UserResponse {
	response.Status = false
	response.ErrorType = Error.Error()
	response.Message = key
	response.Request = request
	return response
}

func (response *UserResponse) Success(ctx *context.Context, key string, data *dto.User) *UserResponse {
	response.Status = true
	response.Message = key
	response.Data = data
	return response
}
