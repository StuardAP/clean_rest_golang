// cmd/main.go
package main

import (
	"log"

	"github.com/StuardAP/clean_rest_golang/pkg/config"
	"github.com/StuardAP/clean_rest_golang/infrastructure/http"
	db "github.com/StuardAP/clean_rest_golang/pkg/database/mysql"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
    db, err := db.SetupMySQLDatabase(mysql.Config{
        User:                 config.Envs.DatabaseConfig.DBUser,
        Passwd:               config.Envs.DatabaseConfig.DBPassword,
        Addr:                 config.Envs.DatabaseConfig.DBAddress,
        DBName:               config.Envs.DatabaseConfig.DBName,
        Net:                  "tcp",
        AllowNativePasswords: true,
        ParseTime:            true,
    })

    if err != nil {
        log.Fatalf("ERROR CREATING DATABASE CONNECTION: %v", err)
    }

    app := fiber.New()

    http.SetupRoutes(app, db)

    log.Printf("STARTING SERVER ON PORT %s", config.Envs.ServerConfig.Port)
    if err := app.Listen(":" + config.Envs.ServerConfig.Port); err != nil {
        log.Fatalf("ERROR STARTING SERVER: %v", err)
    }
}
