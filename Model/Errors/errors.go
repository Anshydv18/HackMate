package hmerrors

import (
	base "Hackmate/Base"
	"context"
	"net/http"
)

type Bderror struct {
	StatusCode   int `json:"status"`
	Status       bool
	Message      string
	Actual_error error
	Data         interface{}
}

func InvalidInputError(ctx *context.Context, key string, request ...interface{}) *Bderror {
	return &Bderror{
		StatusCode: http.StatusBadRequest,
		Status:     false,
		Message:    key,
		Data:       request,
	}
}

func DataBaseConnectionError(ctx *context.Context, key string, request ...interface{}) *Bderror {
	base.Initiate()
	return &Bderror{
		StatusCode: http.StatusForbidden,
		Status:     false,
		Message:    key,
		Data:       request,
	}
}

func DataBaseReadError(ctx *context.Context, error string, request ...interface{}) *Bderror {
	return &Bderror{
		StatusCode: http.StatusForbidden,
		Status:     false,
		Message:    error,
		Data:       request,
	}
}
