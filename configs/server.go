package configs

import (
	"log"
	model "portfolio-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

// StartServer func for starting a simple server.
func StartServer(a *fiber.App, env model.Environment) {
	if err := a.Listen(env.BASE_SERVER_URL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
