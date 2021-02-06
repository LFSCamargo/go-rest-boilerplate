package router

import (
	userRouter "github.com/LFSCamargo/go-rest-boilerplate/modules/user/router"
	"github.com/gofiber/fiber/v2"
)

// Router - All the routes of the application
func Router(app *fiber.App) {
	userRouter.ApplyRouter(app)
}
