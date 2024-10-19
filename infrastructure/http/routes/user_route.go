package routes

import (
	"database/sql"

	queryUser "github.com/StuardAP/clean_rest_golang/infrastructure/database/queries/user"
	handlerUser "github.com/StuardAP/clean_rest_golang/infrastructure/http/handlers/user"
	usecaseUser "github.com/StuardAP/clean_rest_golang/internal/domain/usecases/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router, db *sql.DB) {
    UserRepository := queryUser.NewUserRepositoryMySQL(db)
    userUseCase := usecaseUser.NewUserUseCase(UserRepository)
    userHandler := handlerUser.NewUserHandler(userUseCase)

    userHandler.RegisterRoutes(api)
}
