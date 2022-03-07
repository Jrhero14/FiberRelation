package main

import (
	"github.com/Jrhero14/FiberRelation/database"
	"github.com/Jrhero14/FiberRelation/routes"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	routes.SetupRoutes(app)

	// Listen on PORT 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	errListen := app.Listen(":" + port)
	if errListen != nil {
		os.Exit(1)
	}
}
