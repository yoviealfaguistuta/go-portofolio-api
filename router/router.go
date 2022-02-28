package router

import (
	"portfolio-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupMainRoutes(router fiber.Router) {

	router.Post("/login", controllers.Login)
}
