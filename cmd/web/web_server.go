package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kyp0717/ew-system/controllers"
	"github.com/kyp0717/ew-system/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	if err := godotenv.Load(".env.postgres"); err != nil {
		log.Fatal("Error in loading .env file.")
	}
	controllers.PgConnectDB()

}

func main() {
	app := fiber.New(fiber.Config{
		// Setting centralized error hanling.
		ErrorHandler: handlers.CustomErrorHandler,
	})

	app.Static("/", "./assets")

	app.Use(logger.New())

	handlers.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
