package util

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/fgrosse/goldi"
	"github.com/labstack/echo"
)

type (
	// ResponsePattern return object for json response template
	ResponsePattern struct {
		Status  string      `json:"status"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Meta    interface{} `json:"meta"`
	}

	// CustomApplicationContext return object implementation of custom context and utils methods implementation
	CustomApplicationContext struct {
		echo.Context
		Container  *goldi.Container
		AWSSession *session.Session
	}
)

// CustomResponse return object for json template as custom response
func (c *CustomApplicationContext) CustomResponse(status string, data interface{}, message string, code int, meta interface{}) error {
	return c.JSON(code, &ResponsePattern{
		Status:  status,
		Data:    data,
		Message: message,
		Code:    code,
		Meta:    meta,
	})
}
