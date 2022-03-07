package routes

import (
	userhandler "github.com/Jrhero14/FiberRelation/UserApp/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")
	// Get All User
	user.Get("/", userhandler.GetUser)
	// Create New User
	user.Post("/createUser", userhandler.CreateUser)
}
