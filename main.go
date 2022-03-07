package main

import (
	"FiberRelation/database"
	"FiberRelation/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	routes.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":3000")
}
