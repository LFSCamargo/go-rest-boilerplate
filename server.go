package main

import (
	"github.com/LFSCamargo/go-rest-boilerplate/database"
	"github.com/LFSCamargo/go-rest-boilerplate/env"
	"github.com/LFSCamargo/go-rest-boilerplate/modules/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	config := env.GetConfig()

	router.Router(app)

	if config.Env == "DEV" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		}))
	} else {
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:3000",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))
	}

	database.Connect()

	app.Listen(config.Port)
}
