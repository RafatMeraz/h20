package main

import (
	"fmt"
	"github.com/RafatMeraz/h20/database"
	"github.com/RafatMeraz/h20/routers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	database.Database.Connect()
}

func main() {
	e := echo.New()
	routers.Router{}.RegisterRoutes(e)
	e.Use(middleware.Logger())
	log.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%v", os.Getenv("SERVER_PORT"))))
}
