package routes

import (
	"fmt"

	"github.com/averystampp/boilerplate/gobackend/database"
	model "github.com/averystampp/boilerplate/gobackend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

var store = session.New()

func Login(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(model.User)
	c.BodyParser(user)
	item := model.User{}
	db.First(&item, "username = ?", user.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(item.Password), []byte(user.Password)); err != nil {
		return c.SendString("Sorry cannot login")
	}
	sess, err := store.Get(c)
	if err != nil {
		fmt.Println(err)
	}
	defer sess.Save()
	sess.Regenerate()
	sess.Set("session-token", sess.ID())
	fmt.Println(sess.Get("session-token"))
	return c.JSON(sess.ID())

}
