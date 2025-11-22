package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() fiber.Handler {
	secret := os.Getenv("JWT_SECRET")

	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return c.Status(401).JSON(fiber.Map{"error": "No token"})
		}

		tokenString := auth[len("Bearer "):]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		c.Locals("user", token)
		return c.Next()
	}
}
