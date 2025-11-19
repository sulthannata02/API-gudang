package routes

import (
	"gudang-app/middleware"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"gudang-app/handlers"
)

func Setup(app *fiber.App) {

	// Auth
	app.Post("/api/register", handlers.Register)
	app.Post("/api/login", handlers.Login)

	// JWT middleware untuk semua route berikutnya
	app.Use(jwtware.New(jwtware.Config{
		SigningKey:   []byte("kobra-secret"),
		ContextKey:   "user",
		ErrorHandler: jwtError,
	}))

	api := app.Group("/api")

	// =====================
	// Barang
	// =====================
	// Semua user bisa lihat barang
	api.Get("/barang", handlers.GetBarang)

	// Hanya admin yang bisa CRUD barang
	api.Post("/barang", middleware.RoleMiddleware("admin"), handlers.CreateBarang)
	api.Put("/barang/:id", middleware.RoleMiddleware("admin"), handlers.UpdateBarang)
	api.Delete("/barang/:id", middleware.RoleMiddleware("admin"), handlers.DeleteBarang)

	// =====================
	// Transaksi
	// =====================
	// Admin & Staff bisa transaksi
	api.Get("/transaksi", middleware.RoleMiddleware("admin", "staff"), handlers.GetTransaksi)
	api.Post("/transaksi", middleware.RoleMiddleware("admin", "staff"), handlers.CreateTransaksi)
}

// error handler JWT
func jwtError(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	return nil
}
