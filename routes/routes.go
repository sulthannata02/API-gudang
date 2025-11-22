package routes

import (

	"github.com/gofiber/fiber/v2"
	"gudang-app/handlers"
	"gudang-app/middleware"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", handlers.Register)
	app.Post("/api/login", handlers.Login)

	api := app.Group("/api")

	api.Use(middleware.JWTMiddleware()) // <= GANTI INI

	api.Get("/barang", handlers.GetBarang)
	api.Post("/barang", middleware.RoleMiddleware("admin"), handlers.CreateBarang)
	api.Put("/barang/:id", middleware.RoleMiddleware("admin"), handlers.UpdateBarang)
	api.Delete("/barang/:id", middleware.RoleMiddleware("admin"), handlers.DeleteBarang)

	api.Get("/transaksi", middleware.RoleMiddleware("admin", "staff"), handlers.GetTransaksi)
	api.Post("/transaksi", middleware.RoleMiddleware("admin", "staff"), handlers.CreateTransaksi)
}
