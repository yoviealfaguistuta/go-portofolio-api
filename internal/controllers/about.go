package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AboutControllers struct {
	contextTimeout time.Duration
	dbConn         *pgxpool.Pool
}

func NewAboutControllers(conn *pgxpool.Pool, timeout time.Duration) *AboutControllers {
	return &AboutControllers{
		dbConn:         conn,
		contextTimeout: timeout,
	}
}

func (dc *AboutControllers) About(c *fiber.Ctx) (responses string, err error) {

	query_list := "SELECT messages FROM about WHERE id = 1;"
	row := dc.dbConn.QueryRow(context.Background(), query_list)
	err = row.Scan(&responses)

	if err != nil {
		return
	}

	return
}
