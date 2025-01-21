package hmerrors

import (
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

func InvalidInputError(ctx *context.Context, key string, request interface{}) *Bderror {
	return &Bderror{
		StatusCode: http.StatusBadRequest,
		Status:     false,
		Message:    key,
		Data:       request,
	}
}
