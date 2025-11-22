package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"gudang-app/config"
	"gudang-app/routes"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	app := fiber.New()

	// CORS config (ambil origin dari env)
	frontend := os.Getenv("FRONTEND_ORIGIN")
	if frontend == "" {
		frontend = "http://localhost:5173"
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins:     frontend,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))

	// Database
	config.Connect()
	config.AutoMigrate()

	// Routes
	routes.Setup(app)

	// Run server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("Listening on :" + port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
