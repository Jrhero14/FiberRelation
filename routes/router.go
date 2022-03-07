package routes

import (
	userRoutes "github.com/Jrhero14/FiberRelation/UserApp/routes"
	noteRoutes "github.com/Jrhero14/FiberRelation/noteApp/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Setup the Node Routes
	noteRoutes.SetupNoteRoutes(api)
	// Setup the User Routes
	userRoutes.SetupUserRoutes(api)
}
