package configs

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	model "portfolio-api/internal/models"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

/*
DATABASE CONNECTION
  - Open Database
  - Check Database Connection
  - Run Migration
  - Start Connection
*/
func Database(env model.Environment) (CONNECTION *pgxpool.Pool) {

	var DATABASE *sql.DB
	var DRIVER database.Driver
	var MIGRATION *migrate.Migrate
	var err error

	// Open Database: opens a database specified by its database
	// DRIVER name and a DRIVER-specific data source name
	DATABASE, err = sql.Open("postgres", env.BASE_DB_SERVER_URL)
	if err != nil {
		Disconnect("Failed to Open Database: %v\n", err)
	}

	// Check Database Connection: Make sure database connection
	// is running without problems
	DRIVER, err = postgres.WithInstance(DATABASE, &postgres.Config{})
	if err != nil {
		_ = DATABASE.Close()
		Disconnect("Failed to Check Database Connection: %v\n", err)
	}

	// Run Migrations: Initiate MIGRATION or generate some
	// tables from "migrations" folder
	MIGRATION, err = migrate.NewWithDatabaseInstance("file://migrations", "postgres", DRIVER)
	if err != nil {
		_ = DATABASE.Close()
		Disconnect("Failed to Run Migrations: %v\n", err)
	}

	// *note: Uncomment this to make change tables rows or something likes modify table structure
	// *warning: You must delete all tables and MIGRATION, this will cause all data will be lost
	// err = MIGRATION.Down()
	// if err != nil {
	// 	log.Println("Some Problem With Migration Down:" + err.Error())
	// }

	err = MIGRATION.Up()
	if err != nil {
		log.Println("Some Problem With Migration UP:" + err.Error())
	}

	err = DATABASE.Close()
	if err != nil {
		log.Println("Failed to Close Database Connection:" + err.Error())
	}

	// Start Connection: This will start to connect postgres database
	CONNECTION, err = pgxpool.Connect(context.Background(), env.BASE_DB_SERVER_URL)
	if err != nil {
		Disconnect("Unable to Start Postgres Database: %v\n", err)
	}

	return CONNECTION
}

func Disconnect(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}
