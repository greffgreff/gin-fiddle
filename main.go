package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/greffgreff/gin-fiddle/src/controllers"
	"github.com/greffgreff/gin-fiddle/src/databases"
	"github.com/greffgreff/gin-fiddle/src/middlewares"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databases.ConnectToDB()
}

func main() {
	app := gin.New()

	routerGroup := app.Group("/api/v1")
	controllers.InitializeEmployeeController(routerGroup)

	app.Use(middlewares.ErrorHandler)
	app.Run()
}
