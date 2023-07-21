package routers

import (
	"github.com/RafatMeraz/h20/controllers"
	"github.com/RafatMeraz/h20/middlewares"
	"github.com/labstack/echo/v4"
)

type Router struct{}

func (Router) RegisterRoutes(e *echo.Echo) {
	e.GET("/", controllers.HomeController{}.Home).Name = "Home"
	e.POST("/login", controllers.AuthController{}.Login).Name = "Login"
	e.POST("/signup", controllers.AuthController{}.SignUp).Name = "SignUp"
	authRouteGroup := e.Group("/", middlewares.AuthMiddleware{}.AuthVerification)
	authRouteGroup.GET("water-track/:id", controllers.WaterTrackerController{}.GetUserWaterTrack).Name = "Get water track of user"
	authRouteGroup.GET("water-track", controllers.WaterTrackerController{}.GetUserWaterTrack).Name = "Self water track history"
}
