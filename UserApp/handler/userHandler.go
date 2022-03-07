package handler

import (
	"FiberRelation/database"
	"FiberRelation/noteApp/model"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user []model.User

	// Find All User
	db.Preload("Notes").Find(&user)
	//db.Find(&user)

	// Cek ketersediaan
	if len(user) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Tidak ada User", "data": nil})
	}

	// Kalau ada
	return c.JSON(fiber.Map{"status": "success", "message": "User ditemukan", "data": user})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	// menyerahkan data dari body
	err := c.BodyParser(user)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"status":  "error internal server",
			"message": "goblok",
			"data":    nil,
		})
	}
	// Add user to User Model
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// Return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created User", "data": user})
}
