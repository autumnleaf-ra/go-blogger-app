package routes

import (
	"github.com/autumnleaf-ra/go-blogger-app/controller"
	"github.com/gofiber/fiber/v2"
)

// routes to use api
func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	// app.Use(middleware.IsAuthenticate)
}
