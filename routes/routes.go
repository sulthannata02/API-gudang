package routes

import (
    "github.com/gofiber/fiber/v2"
    "gudang-app/handlers"
    "gudang-app/middleware"
)

func Setup(app *fiber.App) {
    // Route tanpa JWT
    app.Post("/api/register", handlers.Register)
    app.Post("/api/login", handlers.Login)

    // Group API dengan JWT
    api := app.Group("/api", middleware.JWTMiddleware())

    // Barang
    api.Get("/barang", middleware.RoleMiddleware("admin", "staff"), handlers.GetBarang) // ✅ admin & staff bisa lihat
    api.Post("/barang", middleware.RoleMiddleware("admin"), handlers.CreateBarang)      // hanya admin
    api.Put("/barang/:id", middleware.RoleMiddleware("admin"), handlers.UpdateBarang)   // hanya admin
    api.Delete("/barang/:id", middleware.RoleMiddleware("admin"), handlers.DeleteBarang) // hanya admin

    // Transaksi (admin + staff)
    api.Get("/transaksi", middleware.RoleMiddleware("admin", "staff"), handlers.GetTransaksi)
    api.Post("/transaksi", middleware.RoleMiddleware("admin", "staff"), handlers.CreateTransaksi)
}
