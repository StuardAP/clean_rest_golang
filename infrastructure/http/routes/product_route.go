// http/routes/product_routes.go
package routes

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	handlerProduct "github.com/StuardAP/clean_rest_golang/infrastructure/http/handlers/product"
	usecaseProduct "github.com/StuardAP/clean_rest_golang/internal/domain/usecases/product"
	queryProduct "github.com/StuardAP/clean_rest_golang/infrastructure/database/queries/product"
)

func SetupProductRoutes(api fiber.Router, db *sql.DB) {
    ProductRepository := queryProduct.NewProductRepositoryMySQL(db)
    productUseCase := usecaseProduct.NewProductUseCase(ProductRepository)
    productHandler := handlerProduct.NewProductHandler(productUseCase)

    productHandler.RegisterRoutes(api)
}
