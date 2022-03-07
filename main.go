package main

import (
	"github.com/Jrhero14/FiberRelation/database"
	"github.com/Jrhero14/FiberRelation/routes"
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
