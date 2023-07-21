package routers

import (
	"github.com/RafatMeraz/h20/controllers"
	"github.com/labstack/echo/v4"
)

type Router struct{}

func (Router) RegisterRoutes(e *echo.Echo) {
	e.GET("/", controllers.HomeController{}.Home).Name = "Home"
	e.POST("/login", controllers.AuthController{}.Login).Name = "Login"
	e.POST("/signup", controllers.AuthController{}.SignUp).Name = "SignUp"
}
