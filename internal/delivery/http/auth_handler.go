package http

import (
	"boilerplate-go/internal/domain/service"
	"context"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
    service service.AuthService
}


func NewAuthHandler(service service.AuthService) *AuthHandler {
    return &AuthHandler{service: service}
}

func (h *AuthHandler) RegisterRoutes(app *fiber.App) {
    app.Post("/login", h.Login)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var request service.LoginRequest
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
    }

    token, err := h.service.Login(context.Background(), request)
    
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(token)
}