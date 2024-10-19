package handlers

import (
	usecase "github.com/StuardAP/clean_rest_golang/internal/domain/usecases/user"
	"github.com/StuardAP/clean_rest_golang/pkg/config"
	auth "github.com/StuardAP/clean_rest_golang/pkg/utilities/auth"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/register", h.RegisterUser)
	router.Post("/login", h.LoginUser)
	router.Get("/users", h.handleGetAllUsers)
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	type Request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	_, err := h.userUseCase.GetUserByEmail(req.Email)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
	}

	user, err := h.userUseCase.RegisterUser(req.Name, req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register user", "details": err.Error()})
	}

	return c.JSON(user)
}

func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	user, err := h.userUseCase.GetUserByEmail(req.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if err := auth.ComparePassword(user.Password, req.Password); err != nil {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	secret := []byte(config.Envs.JWTConfig.JWTSecret)
	token, err := auth.CreateJWT(secret, user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (h *UserHandler) handleGetAllUsers(c *fiber.Ctx) error {
	users, err := h.userUseCase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No users found"})
	}
	return c.JSON(users)
}
