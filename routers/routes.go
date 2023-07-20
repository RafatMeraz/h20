package routers

import (
	"github.com/RafatMeraz/h20/auth/controllers"
	"github.com/RafatMeraz/h20/home"
	"github.com/labstack/echo/v4"
)

type Router struct{}

func (Router) RegisterRoutes(e *echo.Echo) {
	e.GET("/", home.HomeController{}.Home).Name = "Home"
	e.POST("/login", controllers.AuthController{}.Login).Name = "Login"
	e.POST("/signup", controllers.AuthController{}.SignUp).Name = "SignUp"
}
