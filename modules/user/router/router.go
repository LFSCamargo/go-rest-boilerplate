package userRouter

import (
	"github.com/LFSCamargo/go-rest-boilerplate/middlewares"
	userService "github.com/LFSCamargo/go-rest-boilerplate/modules/user/service"
	"github.com/gofiber/fiber/v2"
)

// ApplyRouter - apply routes for the user
func ApplyRouter(app *fiber.App) {
	app.Post("/login", userService.HandleLogin)
	app.Post("/signup", userService.HandleSignUp)
	app.Get("/me", middlewares.AuthMiddleware, userService.HandleMe)
}
