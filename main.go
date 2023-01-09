package main

import (
	"portfolio-api/configs"
	_ "portfolio-api/docs"
	controllers "portfolio-api/internal/controllers"
	handlers "portfolio-api/internal/handlers/http"
	middlewares "portfolio-api/internal/middlewares"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	env := configs.Init()

	connection := configs.Database(env)
	defer connection.Close()

	timeoutContext := time.Duration(env.BASE_SERVER_READ_TIMEOUT) * time.Second
	fiberConfig := configs.FiberConfig(env.BASE_SERVER_NAME, env.BASE_SERVER_READ_TIMEOUT)
	validator := configs.NewValidator()

	APP := fiber.New(fiberConfig)

	APP.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index.html", "")
	})

	APP.Static("/assets/", "./assets/uploads")
	APP.Get("/swagger/*", swagger.HandlerDefault)

	middL := middlewares.InitMiddleware(APP)
	APP.Use(middL.CORS())
	APP.Use(middL.LOGGER())

	portofolioController := controllers.NewPortofolioControllers(connection, timeoutContext)
	handlers.NewPortofolioHandler(APP, validator, portofolioController)

	aboutController := controllers.NewAboutControllers(connection, timeoutContext)
	handlers.NewAboutHandler(APP, validator, aboutController)

	certificateController := controllers.NewCertificateControllers(connection, timeoutContext)
	handlers.NewCertificateHandler(APP, validator, certificateController)

	configs.StartServer(APP, env)
}
