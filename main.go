package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

var serverName, serverUrl, serverReadTimeout, dbServerUrl string
var db *sql.DB
var err error
var pgxConn *pgxpool.Pool
var driver database.Driver

var migration *migrate.Migrate

func init() {
	// Server Env
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

	// JWT Env
	if os.Getenv("JWT_SECRET_KEY") == "" {
		exitf("JWT_SECRET_KEY env is required")
	}
	if os.Getenv("JWT_EXPIRE_MINUTES") == "" {
		exitf("JWT_EXPIRE_MINUTES env is required")
	}

	// Databse Env
	dbServerUrl = os.Getenv("DB_SERVER_URL")
	if dbServerUrl == "" {
		exitf("DB_SERVER_URL config is required")
	}
}

func main() {
	// database migration
	db, err = sql.Open("postgres", dbServerUrl)
	if err != nil {
		exitf("Db open error: %v\n", err)
	}
	driver, err = postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		_ = db.Close()
		exitf("Db postgres driven error: %v\n", err)
	}
	migration, err = migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		_ = db.Close()
		exitf("Unable to initiate migration: %v\n", err)
	}

	// If you want clear all data rows, uncomment this
	// err = migration.Down() // or m.Step(2) if you want to explicitly set the number of migrations to run
	// if err != nil {
	// 	log.Println(fmt.Sprintf("Migration error: %s", err.Error()))
	// }

	err = migration.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil {
		log.Println(fmt.Sprintf("Migration error: %s", err.Error()))
	}
	err = db.Close()
	if err != nil {
		log.Println(fmt.Sprintf("Db close error: %s", err.Error()))
	}
	// end database migration

	// Pgx Pool Connection
	pgxConn, err = pgxpool.Connect(context.Background(), dbServerUrl)
	if err != nil {
		exitf("Unable to connect to database: %v\n", err)
	}
	defer pgxConn.Close()

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

	// rAuth := app.Group("/auth", middL.JWT()) // router for api private access

	portofolioController := controllers.NewPortofolioControllers(pgxConn, timeoutContext)
	handlers.NewPortofolioHandler(app, validator, portofolioController)

	aboutController := controllers.NewAboutControllers(pgxConn, timeoutContext)
	handlers.NewAboutHandler(app, validator, aboutController)

	configs.StartServer(app)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}
