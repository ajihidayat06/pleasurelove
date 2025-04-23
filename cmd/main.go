package main

import (
	"context"
	"fmt"
	"os/signal"
	"pleasurelove/config"
	"pleasurelove/internal/middleware"
	"pleasurelove/internal/router"
	"pleasurelove/internal/seeder"
	"pleasurelove/internal/utils"
	"pleasurelove/pkg/logger"
	"pleasurelove/pkg/redis"
	"syscall"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting application...")
	fmt.Printf("DEBUG_MODE: %s\n", os.Getenv("DEBUG_MODE"))

	logger.InitLogger()
	logger.Info(context.Background(), "Starting API server...", nil)

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or unable to load it")
	}

	cfg := config.LoadConfig()
	db := config.InitDatabase(cfg)
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	config.RunMigrations()

	err = redis.InitRedis()
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}

	if err := seeder.SeedSuperAdmin(context.Background(), db); err != nil {
		logger.Error(context.Background(), "Failed to seed superadmin", err)
		log.Fatalf("Failed to seed superadmin: %v", err)
	}

	// Security middleware: Helmet untuk secure headers
	setupMiddlewares(app)

	utils.InitValidator()
	router.SetupRoutes(app, db)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   "Route not found",
			"message": "The requested endpoint does not exist",
		})
	})

	port := getPort()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		if err := app.Shutdown(); err != nil {
			log.Fatalf("Error shutting down server: %v", err)
		}
	}()

	log.Printf("Starting server on port %s...", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped.")
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080" // Default port
	}
	return port
}

func setupMiddlewares(app *fiber.App) {
	app.Use(helmet.New()) // Secure headers
	app.Use(config.CorsConfig())
	app.Use(middleware.SetTraceIDAndRequestIDMiddleware) // Fiber context to standard context
	app.Use(middleware.LoggingMiddleware)                // Logging
	app.Use(middleware.RecoverMiddleware())              // Recovery
}
