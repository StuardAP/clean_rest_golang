package handlers

import (
	usecase "github.com/StuardAP/clean_rest_golang/internal/domain/usecases/product"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productUseCase *usecase.ProductUseCase
}

func NewProductHandler(productUseCase *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase: productUseCase}
}

func (h *ProductHandler) RegisterRoutes(router fiber.Router) {

	router.Get("/products", h.handleGetAllProducts)
}



func (h *ProductHandler) handleGetAllProducts(c *fiber.Ctx) error {
	products, err := h.productUseCase.GetAllProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if len(products) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No products found"})
	}
	return c.JSON(products)
}
