package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Golang is Starting")
	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	// router.SetupMainRoutes(app)

	app.Get("/", RoutDefault)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file, %s", err.Error())
	}

	app.Listen(os.Getenv("BASE_URL"))
}

func RoutDefault(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"data":    "OK",
		"message": "Success",
		"code":    200,
		"_func":   "Login",
	})

}
