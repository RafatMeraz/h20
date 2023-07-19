package main

import (
	"fmt"
	"github.com/RafatMeraz/h20/config"
	"github.com/RafatMeraz/h20/database"
	"github.com/RafatMeraz/h20/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func init() {
	database.Database.Connect()
}

func main() {
	e := echo.New()
	routers.Router{}.RegisterRoutes(e)
	e.Use(middleware.Logger())
	log.Fatal(e.Start(fmt.Sprintf("localhost:%v", config.AppConfiguration.Port)))
}
