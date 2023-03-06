package main

import (
	"log"
	"os"

	"github.com/Chucky22Mendoza/Rest-api/repositories"
	"github.com/Chucky22Mendoza/Rest-api/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// database connection
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Connection with database was unsuccesfull")
	} else {
		log.Println("Connection with database succesfully")
	}

	// Declare routes/repositories
	tasksRepository := repositories.Task{
		DB: db,
	}

	// Server init
	app := fiber.New(fiber.Config{
		BodyLimit: 500 * 1024 * 1024,
	})

	// Middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger.New())

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")
	tasksRepository.SetUpRoutes(api)

	// Server listen
	app.Listen(os.Getenv("SERVER_PORT"))
}
