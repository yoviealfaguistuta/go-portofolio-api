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

type PortofolioControllers struct {
	contextTimeout time.Duration
	dbConn         *pgxpool.Pool
}

func NewPortofolioControllers(conn *pgxpool.Pool, timeout time.Duration) *PortofolioControllers {
	return &PortofolioControllers{
		dbConn:         conn,
		contextTimeout: timeout,
	}
}

func (dc *PortofolioControllers) List(c *fiber.Ctx, page int, limit int) (responses []models.PortofolioList, page_count float64, total int, offset int, err error) {

	offset = limit * (page - 1)

	SQL := `SELECT 
				COUNT(*) OVER() as total, 
				p.id,
				p.title, 
				p.descriptions, 
				p.created_at, 
				pi.images 
			FROM portofolio AS p
				LEFT JOIN LATERAL (
					SELECT 
						pi.images
					FROM p_images AS pi
					WHERE p.id = pi.id_portfolio LIMIT 1) pi ON true 
			ORDER BY p.id DESC LIMIT $1 OFFSET $2`

	rows, err := dc.dbConn.Query(context.Background(), SQL, limit, offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var model = new(models.PortofolioList)
		err = rows.Scan(
			&total,
			&model.ID,
			&model.Title,
			&model.Descriptions,
			&model.CreatedAt,
			&model.Images,
		)

		if err != nil {
			return
		}
		responses = append(responses, *model)
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

func (dc *PortofolioControllers) Detail(c *fiber.Ctx, id int) (responses models.PortofolioDetail, err error) {

	SQL := `SELECT 
		p.id, 
		p.title, 
		p.descriptions, 
		p.project_info, 
		p.languages, 
		p.tech, 
		p.job_role, 
		p.databases, 
		p.dates, 
		p.platform, 
		p.urls, 
		p.source_code, 
		p.created_at, 
		p.updated_at, 
		(SELECT json_agg(t) FROM (SELECT p_images.orders, p_images.images FROM p_images WHERE p_images.id_portfolio=p.id) t) AS childs 
	FROM portofolio AS p
	WHERE p.id = $1`

	var images []map[string]interface{}
	err = dc.dbConn.QueryRow(context.Background(), SQL, id).Scan(
		&responses.ID,
		&responses.Title,
		&responses.Descriptions,
		&responses.ProjectInfo,
		&responses.Languages,
		&responses.Tech,
		&responses.JobRole,
		&responses.Databases,
		&responses.Dates,
		&responses.Platform,
		&responses.Urls,
		&responses.SourceCode,
		&responses.CreatedAt,
		&responses.UpdatedAt,
		&images,
	)

	for _, v := range images {
		pi := models.PortofolioImages{
			Orders: int64(v["orders"].(float64)),
			Images: v["images"].(string),
		}
		responses.Images = append(responses.Images, pi)
	}

	if err != nil {
		return
	}

	return
}
