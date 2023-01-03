package routes

import (
	auth "github.com/averystampp/boilerplate/gobackend/authentication"
	"github.com/gofiber/fiber/v2"
)

func Authroute(c *fiber.Ctx) error {
	if auth.CheckID(c) {
		return c.JSON("Auth works")

	}

	return c.JSON("Failed to auth")

}
