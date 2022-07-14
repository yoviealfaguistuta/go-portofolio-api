package main

import (
	"fmt"
	"os"
	"portfolio-api/configs"
	_ "portfolio-api/docs" // load API Docs files (Swagger)
	controllers "portfolio-api/internal/controllers"
	handlers "portfolio-api/internal/handlers"
	middlewares "portfolio-api/internal/middlewares"
	"strconv"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var serverName, serverUrl, serverReadTimeout string
var err error
var pgxConn *pgxpool.Pool

func init() {
	serverName = os.Getenv("SERVER_NAME")
	if serverName == "" {
		exitf("SERVER_NAME env is required")
	}
	serverUrl = os.Getenv("SERVER_URL")
	if serverUrl == "" {
		exitf("SERVER_URL env is required")
	}
	serverReadTimeout = os.Getenv("SERVER_READ_TIMEOUT")
	if serverReadTimeout == "" {
		exitf("SERVER_READ_TIMEOUT env is required")
	}
}
func main() {

	config := configs.FiberConfig()
	app := fiber.New(config)

	app.Static("/assets", "./assets")

	// Swagger handler
	app.Get("/swagger/*", swagger.HandlerDefault)

	middL := middlewares.InitMiddleware(app)
	app.Use(middL.CORS())
	app.Use(middL.LOGGER())

	serverReadTimeoutInt, err := strconv.Atoi(serverReadTimeout)
	if err != nil {
		exitf("Failed casting timeout context: ", err)
	}
	timeoutContext := time.Duration(serverReadTimeoutInt) * time.Second

	validator := configs.NewValidator()

	rAuth := app.Group("/auth", middL.JWT()) // router for api private access

	portofolioController := controllers.NewPortofolioControllers(pgxConn, timeoutContext)
	handlers.NewPortofolioHandler(rAuth, validator, portofolioController)

	configs.StartServer(app)

	// fmt.Println("Golang is Starting")
	// app := fiber.New()
	// app.Use(cors.New())
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))
	// routers.SetupMainRoutes(app)

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file, %s", err.Error())
	// }

	// app.Listen(os.Getenv("BASE_URL"))
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}
