package main

import (
	"demo/go-fiber/config"
	"demo/go-fiber/internal/home"
	"demo/go-fiber/internal/vacancy"
	"demo/go-fiber/pkg/database"
	"demo/go-fiber/pkg/logger"
	"demo/go-fiber/pkg/middleware"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
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

	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})

	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))

	// Repositories
	vacancyRepo := vacancy.NewRepository(dbpool, customLogger)

	// Handlers
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	home.NewHandler(home.HomeHandlerDeps{
		Router:       app,
		CustomLogger: customLogger,
		Repository:   vacancyRepo,
		Store:        store,
	})

	app.Listen(":3000")
}
