package routes

import (
	"fmt"

	"github.com/averystampp/boilerplate/gobackend/database"
	model "github.com/averystampp/boilerplate/gobackend/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	c.Accepts("application/json")
	db := database.DBConn
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(503)
	}
	pass := user.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Could not register user")
	}
	user.Password = string(hash)
	db.Create(&user)
	return c.SendString("User " + user.Username + " has been created")
}
