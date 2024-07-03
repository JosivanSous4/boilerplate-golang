package http

import (
	"boilerplate-go/internal/domain/model"
	"boilerplate-go/internal/domain/service"
	"context"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
    service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
    return &ProductHandler{service: service}
}

func (h *ProductHandler) RegisterRoutes(app *fiber.App) {
    app.Post("/products", h.CreateProduct)
    app.Get("/products/:id", h.GetProductByID)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
    var product model.Product
    if err := c.BodyParser(&product); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse json"})
    }

    err := h.service.CreateProduct(context.Background(), &product)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
    id := c.Params("id")
    product, err := h.service.GetProductByID(context.Background(), id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
    }

    return c.JSON(product)
}
