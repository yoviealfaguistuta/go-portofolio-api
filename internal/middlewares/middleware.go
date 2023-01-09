package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	appCtx *fiber.App
	// another stuff , may be needed by middleware
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "*",
		//AllowOrigins: "http://yoviealfaguistuta.site",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
	})
}

// LOGGER simple logger.
func (m *GoMiddleware) LOGGER() fiber.Handler {
	return logger.New()
}

// InitMiddleware initialize the middleware
func InitMiddleware(ctx *fiber.App) *GoMiddleware {
	return &GoMiddleware{appCtx: ctx}
}
