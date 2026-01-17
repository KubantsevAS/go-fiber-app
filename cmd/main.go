package main

import (
	"demo/go-fiber/config"
	"demo/go-fiber/internal/home"
	"demo/go-fiber/internal/vacancy"
	"demo/go-fiber/pkg/database"
	"demo/go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger((logConfig))

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/static", "./static")

	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()

	// Repositories
	vacancyRepo := vacancy.NewRepository(dbpool, customLogger)

	// Handlers
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	home.NewHandler(app, customLogger)

	app.Listen(":3000")
}
