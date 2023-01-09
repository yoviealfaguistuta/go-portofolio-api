package controllers

import (
	"context"
	"math"
	"portfolio-api/configs"
	"portfolio-api/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
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

func (cc *CertificateControllers) List(c *fiber.Ctx, page int, limit int) (responses []models.Certificate, page_count float64, total int, offset int, err error) {

	offset = limit * (page - 1)

	SQL := `SELECT 
		COUNT(*) OVER() as total, 
		crtcs.id, 
		crtcs.images, 
		crtcs.title, 
		crtcs.publish, 
		crtcs.organization, 
		crtcs.credentials, 
		crtcs.urls 
	FROM certificates AS crtcs ORDER BY crtcs.id DESC LIMIT $1 OFFSET $2`

	rows, err := cc.dbConn.Query(context.Background(), SQL, limit, offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var model = new(models.Certificate)
		err = rows.Scan(
			&total,
			&model.ID,
			&model.Images,
			&model.Title,
			&model.Publish,
			&model.Organization,
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

	page_count = configs.IntToFloat64(total / limit)

	if page_count < 1 {
		page_count = 0
	} else {
		if total%limit == 0 {
			page_count = math.Ceil(page_count)
		} else {
			page_count = math.Ceil(page_count) + 1
		}
	}

	return
}
