package main

import (
	"myapi/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handlers.HelloWorld)

	// New routes for Opportunity
	e.GET("/opportunities/:id", handlers.GetOpportunity)
	e.POST("/opportunities/:id/files", handlers.AddFilesToOpportunity)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
