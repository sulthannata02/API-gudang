package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware cek role
func RoleMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userLoc := c.Locals("user")
		if userLoc == nil {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}

		userToken, ok := userLoc.(*jwt.Token)
		if !ok {
			return c.Status(500).JSON(fiber.Map{"error": "Token type invalid"})
		}

		claims := userToken.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		println("Role dari token:", role) // lihat di console

		for _, allowed := range allowedRoles {
			if role == allowed {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{"error": "Forbidden: role tidak memiliki akses"})
	}
}
