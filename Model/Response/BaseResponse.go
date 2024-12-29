package response

import (
	"context"
)

type BaseResponse struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message ,omitempty"`
	ErrorType string      `json:"error_type ,omitempty"`
	Request   interface{} `json:"request"`
}

func (r *BaseResponse) Fail(ctx *context.Context, key string, error string, request interface{}) *BaseResponse {
	r.Status = false
	r.Message = error
	r.Request = request
	return r
}

func (r *BaseResponse) Success(ctx *context.Context, key string, request interface{}) *BaseResponse {
	r.Status = true
	r.Message = "Success"
	return r
}
