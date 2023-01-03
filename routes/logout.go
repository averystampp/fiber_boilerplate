package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
	}

	if err := sess.Destroy(); err != nil {
		fmt.Println(err)
	}
	return c.SendStatus(200)
}
