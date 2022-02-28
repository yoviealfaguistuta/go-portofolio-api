package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	fmt.Println("Hai")
	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Listen(os.Getenv("BASE_URL"))
}
