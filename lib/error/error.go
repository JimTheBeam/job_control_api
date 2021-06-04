package error

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTPError is a basic HTTP error response
type HTTPError struct {
	Code    int    `json:"status"`
	Name    string `json:"error_name"`
	Message string `json:"message"`
}

// Error implements a custom echo error handler that will encode errors as JSON objects rather than
// just return a text body. It will also make sure to not have the redundant information given in
// the echo string encoding of HTTP errors.
func Error(err error, ctx echo.Context) {
	errObj := HTTPError{
		Code:    http.StatusInternalServerError,
		Name:    http.StatusText(500),
		Message: err.Error(),
	}

	he, ok := err.(*echo.HTTPError)
	if ok {
		errObj.Code = he.Code
		errObj.Message = fmt.Sprintf("%v", he.Message)
		log.Printf("HTTP error code: %d message: %v", he.Code, he.Message)
	}

	errObj.Name = http.StatusText(errObj.Code)
	if !ctx.Response().Committed {
		if ctx.Request().Method == echo.HEAD {
			ctx.NoContent(errObj.Code)
		} else {
			ctx.JSON(errObj.Code, errObj)
		}
	}
}
