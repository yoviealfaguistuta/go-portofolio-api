package controllers

import (
	"context"
	"portfolio-api/internal/models"
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

func (dc *AboutControllers) About(c *fiber.Ctx) (responses models.About, err error) {

	var model = new(models.About)
	query_list := "SELECT messages FROM about WHERE id = 1;"
	row := dc.dbConn.QueryRow(context.Background(), query_list)
	err = row.Scan(&model.Messages)

	responses = *model

	if err != nil {
		return
	}

	return
}
