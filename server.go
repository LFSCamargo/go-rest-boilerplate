package main

import (
	"log"

	"github.com/LFSCamargo/go-rest-boilerplate/database"
	"github.com/LFSCamargo/go-rest-boilerplate/env"
	"github.com/LFSCamargo/go-rest-boilerplate/modules/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config := env.GetConfig()

	router.Router(app)

	database.Connect()

	app.Listen(config.Port)
	log.Printf("Server exposed at localhost:3000")
}
