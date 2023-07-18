package home

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HomeController struct{}

func (HomeController) Home(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"version": "0.0.1",
		"name":    "h20",
	})
}
