package error_mapper

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ErrorMapper struct{}

func (ErrorMapper) MapError(c echo.Context, err error) error {
	e := errorList[err]
	if e.ErrorCode == 0 {
		return err
	}
	if e.Message == "" {
		return c.NoContent(e.ErrorCode)
	}
	return c.JSON(e.ErrorCode, map[string]string{
		"message": e.Message,
	})
}

var errorList = map[error]errorWrapper{
	echo.ErrUnauthorized:   {ErrorCode: 401, Message: "Unauthorized"},
	gorm.ErrRecordNotFound: {ErrorCode: 401, Message: "record not found"},
}

type errorWrapper struct {
	ErrorCode int
	Message   string
}
