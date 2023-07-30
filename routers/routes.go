package routers

import (
	"github.com/RafatMeraz/h20/controllers"
	"github.com/RafatMeraz/h20/middlewares"
	"github.com/RafatMeraz/h20/repositories"
	"github.com/labstack/echo/v4"
)

type Router struct{}

func (Router) RegisterRoutes(e *echo.Echo) {
	authController := controllers.AuthController{
		Repository: repositories.MySqlUserRepository{},
	}
	waterTrackerController := controllers.WaterTrackerController{
		WaterTrackRepository: repositories.MySqlWaterTrackerRepository{},
	}

	e.GET("/", controllers.HomeController{}.Home).Name = "Home"
	// auth routes
	e.POST("/login", authController.Login).Name = "Login"
	e.POST("/signup", authController.SignUp).Name = "SignUp"

	// logged in state routes
	authRouteGroup := e.Group("/", middlewares.AuthMiddleware{}.AuthVerification)
	// water track routes
	authRouteGroup.GET("water-track/:id", waterTrackerController.GetUserWaterTrack).Name = "Get water track of user"
	authRouteGroup.GET("water-track", waterTrackerController.GetUserWaterTrack).Name = "Self water track history"
	authRouteGroup.POST("water-track", waterTrackerController.AddNewWaterTrack).Name = "Add new water track"
	authRouteGroup.DELETE("water-track/:id", waterTrackerController.DeleteWaterTrack).Name = "Delete water track"
}
