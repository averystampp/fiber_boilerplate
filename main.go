package main

import (
	"fmt"

	"github.com/averystampp/firstapp/database"
	post "github.com/averystampp/firstapp/models"
	"github.com/averystampp/firstapp/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initdata() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
	fmt.Println("Connected to DB")
	database.DBConn.AutoMigrate(post.Post{}, post.User{})
	fmt.Println("Database Migrated")
}

func main() {
	// app and app settings/config
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin, mysessioncookie, Authentication, Cookie",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// Setup DB and close it when done running
	initdata()
	// route groupers and versions
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// api/v1 routes
	v1.Get("/", routes.Home)
	v1.Get("/posts", routes.GetPosts)
	v1.Get("/post/:id", routes.GetPost)
	v1.Delete("/post/:id", routes.DeletePost)
	v1.Post("/post/", routes.CreatePost)
	v1.Put("/post/:id", routes.UpdatePost)

	// user routes
	v1.Post("/login", routes.Login)
	v1.Post("/register", routes.Register)
	v1.Delete("/user/:id", routes.DeleteUser)
	v1.Get("/users", routes.AllUsers)
	v1.Get("/logout", routes.Logout)

	// debug routes
	v1.Post("debug", routes.Debug)
	// Port that server is exposed on
	app.Listen(":3000")
}
