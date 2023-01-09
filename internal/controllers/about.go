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

func (ac *AboutControllers) About(c *fiber.Ctx) (responses models.About, err error) {

	SQL := `SELECT messages FROM about WHERE id = 1`
	err = ac.dbConn.QueryRow(context.Background(), SQL).Scan(&responses.Messages)

	if err != nil {
		return
	}

	return
}
