package main

import (
	"fmt"
	"log"
	"os"
	router "portfolio-api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Hai")
	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.SetupMainRoutes(app)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app.Listen(os.Getenv("BASE_URL"))
}
