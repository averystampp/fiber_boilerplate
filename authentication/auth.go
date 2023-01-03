package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type SessionId struct {
	SessionId string `reqHeader:"mysessioncookie"`
}

var store = session.New()

func CheckID(c *fiber.Ctx) bool {
	s := new(SessionId)
	if err := c.ReqHeaderParser(s); err != nil {
		return false
	}

	sess, err := store.Get(c)
	if err != nil {
		return false
	}

	if s.SessionId != sess.ID() {
		fmt.Println("Dont work")
		fmt.Println(s.SessionId)
		fmt.Println(sess.ID())
		return false

	}
	return true
}
