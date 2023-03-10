package routes

import (
	auth "github.com/averystampp/boilerplate/gobackend/authentication"

	"github.com/averystampp/boilerplate/gobackend/database"
	model "github.com/averystampp/boilerplate/gobackend/models"
	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	if auth.CheckID(c) {
		db := database.DBConn
		userId := c.Params("id")
		var user model.User
		db.Delete(&user, userId)
		return c.SendString("Deleted User " + userId)

	}
	return c.SendString("Could not delete user")
}
