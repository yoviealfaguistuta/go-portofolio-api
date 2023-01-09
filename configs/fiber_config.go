package configs

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig(BASE_SERVER_NAME string, BASE_SERVER_READ_TIMEOUT int) fiber.Config {
	return fiber.Config{
		AppName:      BASE_SERVER_NAME,
		ReadTimeout:  time.Second * time.Duration(BASE_SERVER_READ_TIMEOUT),
		ServerHeader: strconv.Itoa(BASE_SERVER_READ_TIMEOUT),
		BodyLimit:    6 * 1024 * 1024,
	}
}
