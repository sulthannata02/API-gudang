package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// cek role dari token yang disimpan jwtware
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

		claims, ok := userToken.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(500).JSON(fiber.Map{"error": "Cannot parse claims"})
		}

		role, ok := claims["role"].(string)
		if !ok {
			return c.Status(500).JSON(fiber.Map{"error": "Role tidak ditemukan di token"})
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{"error": "Forbidden: role tidak memiliki akses"})
	}
}
