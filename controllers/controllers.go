package controllers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func createConnection() *sql.DB {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	// log.Fatalf("Error loading .env file")
	// }

	db, err := sql.Open("postgres", "user=postgres password=localhost dbname=portfolio sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}

func Login(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"data":    "OK",
		"message": "Success",
		"code":    200,
		"_func":   "Login",
	})

}
