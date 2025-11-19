package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gudang-app/config"
	"gudang-app/models"
)

func GetBarang(c *fiber.Ctx) error {
	var barang []models.Barang
	config.DB.Find(&barang)
	return c.JSON(barang)
}

func CreateBarang(c *fiber.Ctx) error {
	var data models.Barang

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Create(&data)
	return c.JSON(data)
}

func UpdateBarang(c *fiber.Ctx) error {
	id := c.Params("id")
	var barang models.Barang

	if err := config.DB.First(&barang, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Barang tidak ditemukan"})
	}

	if err := c.BodyParser(&barang); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	config.DB.Save(&barang)
	return c.JSON(barang)
}

func DeleteBarang(c *fiber.Ctx) error {
	id := c.Params("id")
	var barang models.Barang

	if err := config.DB.Delete(&barang, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Barang tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"message": "Barang berhasil dihapus"})
}
