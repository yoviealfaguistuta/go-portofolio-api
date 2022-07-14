package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		AppName:     os.Getenv("APP_NAME"),
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		//Prefork:       true,
		//CaseSensitive: true,
		//StrictRouting: true,
		ServerHeader: os.Getenv("SERVER_NAME"),
	}
}
