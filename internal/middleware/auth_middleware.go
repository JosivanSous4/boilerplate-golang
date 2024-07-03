package middleware

import (
	"boilerplate-go/internal/infrastructure/security"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        token := c.Get("Authorization")
        if token == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Missing or invalid token",
            })
        }

        claims, err := security.VerifyJWT(token)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Invalid token",
            })
        }

        c.Locals("username", claims["username"])

        return c.Next()
    }
}
