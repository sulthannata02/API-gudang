package handlers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"gudang-app/config"
	"gudang-app/models"
)

// Ambil JWT_SECRET setelah .env diload
func GetJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET belum di-load! Pastikan .env terbaca")
	}
	return []byte(secret)
}

// Register user
func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Username sudah ada"})
	}

	return c.JSON(fiber.Map{"message": "Register berhasil"})
}

// Login user
func Login(c *fiber.Ctx) error {
	req := new(models.User)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Password salah"})
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(GetJWTSecret())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal generate token"})
	}

	return c.JSON(fiber.Map{
		"token": t,
		"role":  user.Role,
	})
}
