package main

import (

	"github.com/gofiber/fiber/v2"	
	"github.com/NotNoss/Go-Web-App-Boilerplate/controllers"
)

func Routes(app *fiber.App) {
    app.Get("/", controllers.PostsIndex)
}