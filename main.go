package main

import (
	"fmt"

	"github.com/averystampp/boilerplate/gobackend/database"
	model "github.com/averystampp/boilerplate/gobackend/models"
	"github.com/averystampp/boilerplate/gobackend/routes"
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
	database.DBConn.AutoMigrate(model.User{})
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

	// route groups and versions
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// api/v1 routes
	v1.Get("/", routes.Home)

	// user routes
	app.Post("/login", routes.Login)
	app.Post("/register", routes.Register)
	app.Delete("/user/:id", routes.DeleteUser)
	app.Get("/logout", routes.Logout)
	app.Get("/auth", routes.Authroute)

	// Port that server is exposed on
	app.Listen(":3000")
}
