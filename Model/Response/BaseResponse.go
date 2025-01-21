package response

import (
	errors "Hackmate/Model/Errors"
	"context"
)

type BaseResponse struct {
	Status  bool            `json:"status"`
	Message string          `json:"message ,omitempty"`
	Error   *errors.Bderror `json:"error_type ,omitempty"`
	Request interface{}     `json:"request"`
}

func (r *BaseResponse) Fail(ctx *context.Context, key string, error *errors.Bderror, request interface{}) *BaseResponse {
	r.Status = false
	r.Message = error.Message
	r.Request = request
	r.Error = error
	return r
}

func (r *BaseResponse) Success(ctx *context.Context, key string, request interface{}) *BaseResponse {
	r.Status = true
	r.Message = "Success"
	return r
}
