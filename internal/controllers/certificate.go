package controllers

import (
	"context"
	"portfolio-api/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CertificateControllers struct {
	contextTimeout time.Duration
	dbConn         *pgxpool.Pool
}

func NewCertificateControllers(conn *pgxpool.Pool, timeout time.Duration) *CertificateControllers {
	return &CertificateControllers{
		dbConn:         conn,
		contextTimeout: timeout,
	}
}

func (dc *CertificateControllers) List(c *fiber.Ctx) (responses []models.Certificate, err error) {
	query := "SELECT id, images, title, publish, credentials, urls FROM certificates;"
	var rows pgx.Rows
	rows, err = dc.dbConn.Query(context.Background(), query)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var model = new(models.Certificate)
		err = rows.Scan(
			&model.ID,
			&model.Images,
			&model.Title,
			&model.Publish,
			&model.Credentials,
			&model.Urls,
		)

		if err != nil {
			return
		}
		responses = append(responses, *model)
	}

	if err != nil {
		return
	}

	return
}
