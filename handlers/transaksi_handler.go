package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gudang-app/config"
	"gudang-app/models"
)

// GET semua transaksi
func GetTransaksi(c *fiber.Ctx) error {
	var transaksi []models.Transaksi
	config.DB.Find(&transaksi)
	return c.JSON(transaksi)
}

// Buat transaksi masuk/keluar
func CreateTransaksi(c *fiber.Ctx) error {
	var req models.Transaksi
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Ambil barang terkait
	var barang models.Barang
	if err := config.DB.First(&barang, req.IDBarang).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Barang tidak ditemukan"})
	}

	// Update stok
	if req.Jenis == "masuk" {
		barang.Stok += req.Jumlah
	} else if req.Jenis == "keluar" {
		if barang.Stok < req.Jumlah {
			return c.Status(400).JSON(fiber.Map{"error": "Stok tidak cukup"})
		}
		barang.Stok -= req.Jumlah
	} else {
		return c.Status(400).JSON(fiber.Map{"error": "Jenis harus 'masuk' atau 'keluar'"})
	}

	// Simpan transaksi dan update stok barang
	config.DB.Save(&barang)
	config.DB.Create(&req)

	return c.JSON(req)
}
