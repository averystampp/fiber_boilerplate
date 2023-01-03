package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Logout(c *fiber.Ctx) error {
	store := session.New()
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
	}

	if err := sess.Destroy(); err != nil {
		fmt.Println(err)
	}
	return c.SendStatus(200)
}
