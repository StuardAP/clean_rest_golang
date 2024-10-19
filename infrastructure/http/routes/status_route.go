package routes

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func SetupStatusRoutes(api fiber.Router, db *sql.DB) {

    api.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status": "success",
            "message": "API is up and running",
        })
    })

    api.Get("/db-check", func(c *fiber.Ctx) error {
        if err := db.Ping(); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "status": "error",
                "message": "Database connection failed",
            })
        }
        return c.JSON(fiber.Map{
            "status": "success",
            "message": "Database is connected",
        })
    })
}
