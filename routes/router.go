package routes

import (
	"fmt"
	"net/http"

	"github.com/averystampp/firstapp/database"
	post "github.com/averystampp/firstapp/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

type SessionId struct {
	SessionId string `reqHeader:"mysessioncookie"`
}

var store = session.New()

func checkID(c *fiber.Ctx) bool {
	s := new(SessionId)
	if err := c.ReqHeaderParser(s); err != nil {
		return false
	}

	sess, err := store.Get(c)
	if err != nil {
		return false
	}
	token := sess.Get("session-token")
	if s.SessionId != token {
		fmt.Println("Dont work")

		return false

	}
	return true
}
func Home(c *fiber.Ctx) error {
	// if !checkID(c) {
	// 	return c.SendStatus(400)
	// }
	return c.JSON("This worked")
}

func GetPosts(c *fiber.Ctx) error {
	c.Accepts("Access-Control-Allow-Origin")
	db := database.DBConn
	var posts []post.Post
	db.Find(&posts)
	return c.JSON(posts)
}

func GetPost(c *fiber.Ctx) error {
	db := database.DBConn
	postId := c.Params("id")
	var post post.Post
	db.Find(&post, postId)
	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	if !checkID(c) {
		return c.SendStatus(400)
	}
	db := database.DBConn
	postId := c.Params("id")
	var post post.Post
	db.Delete(&post, postId)
	return c.SendString("Deleted post number " + postId)
}

func CreatePost(c *fiber.Ctx) error {
	c.Accepts("application/json")
	if !checkID(c) {
		return c.JSON("Bad header")
	}
	db := database.DBConn
	post := new(post.Post)
	if err := c.BodyParser(post); err != nil {
		fmt.Println("got stuck")
		fmt.Println(err)

		return c.SendStatus(503)
	}
	fmt.Println("Got to post stage")
	db.Create(&post)
	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	c.Accepts("application/json")
	if !checkID(c) {
		return c.SendStatus(400)
	}
	postId := c.Params("id")
	if postId == "" {
		return c.SendStatus(400)
	}
	db := database.DBConn
	post := new(post.Post)
	if err := c.BodyParser(post); err != nil {
		return c.SendStatus(503)
	}
	err := db.Model(&post).Where("id=?", postId).Updates(post)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "cannot update post",
		})
	}
	return c.JSON(&fiber.Map{
		"message": "updated post " + postId,
	})

}
func Login(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(post.User)
	c.BodyParser(user)
	item := post.User{}
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

func Register(c *fiber.Ctx) error {
	c.Accepts("application/json")
	db := database.DBConn
	user := new(post.User)
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

func DeleteUser(c *fiber.Ctx) error {
	db := database.DBConn
	userId := c.Params("id")
	var user post.User
	db.Delete(&user, userId)
	return c.SendString("Deleted post number " + userId)
}

func AllUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []post.User
	db.Find(&users)
	return c.JSON(users)
}

func Debug(c *fiber.Ctx) error {
	s := new(SessionId)
	if err := c.ReqHeaderParser(s); err != nil {
		return c.SendString("Cannot get header")
	}

	sess, err := store.Get(c)
	if err != nil {
		return c.SendString("could not get headers")
	}
	token := sess.Get("session-token")
	if s.SessionId != token {
		fmt.Println("Dont work")

		return c.SendString("does not work")

	}
	fmt.Println(sess.ID())
	fmt.Println(s.SessionId)
	fmt.Println(sess.Get("session-token"))
	fmt.Println("work")
	return c.SendString("Checks out")
}
