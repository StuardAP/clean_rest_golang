package http

import (
	"database/sql"

	"github.com/StuardAP/clean_rest_golang/infrastructure/http/routes"
	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App, db *sql.DB) {
    api := app.Group("/api/v1")

    routes.SetupStatusRoutes(api, db)
    routes.SetupUserRoutes(api, db)
    routes.SetupProductRoutes(api, db)
}
