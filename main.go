package main

import (
	"os"

	"github.com/NotNoss/Go-Web-App-Boilerplate/initializers"
	"github.com/NotNoss/Go-Web-App-Boilerplate/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	// Load templates
	engine := html.New("./views", ".html")

	// Setup app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")
	app.Use(middleware.RequireAuth)

	// Routes
	Routes(app)

	// Start app
    app.Listen(":" + os.Getenv("PORT"))
}