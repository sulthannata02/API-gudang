package main

import (
	"github.com/gofiber/fiber/v2"
	"gudang-app/config"
	"gudang-app/routes"
)

func main() {
	app := fiber.New()

	config.Connect()
	config.AutoMigrate()

	routes.Setup(app)

	app.Listen(":3000")
}
