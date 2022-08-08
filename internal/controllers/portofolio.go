package controllers

import (
	"context"
	"portfolio-api/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
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

func (dc *PortofolioControllers) List(c *fiber.Ctx) (responses []models.PortofolioList, err error) {
	query := "SELECT DISTINCT on (portofolio.id) portofolio.id, portofolio.title, portofolio.descriptions, portofolio.created_at, p_images.images FROM portofolio INNER JOIN p_images ON portofolio.id = p_images.id_portfolio ORDER BY portofolio.id, p_images.id ASC"
	var rows pgx.Rows
	rows, err = dc.dbConn.Query(context.Background(), query)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var model = new(models.PortofolioList)
		err = rows.Scan(
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

	if rows.Err() != nil {
		return responses, rows.Err()
	}

	return
}

func (dc *PortofolioControllers) Detail(c *fiber.Ctx, id int) (responses models.PortofolioDetail, err error) {
	var model = new(models.PortofolioDetail)
	var childs []map[string]interface{}
	query := "SELECT portofolio.id, portofolio.title, portofolio.descriptions, portofolio.project_info, portofolio.languages, portofolio.tech, portofolio.job_role, portofolio.databases, portofolio.dates, portofolio.platform, portofolio.urls, portofolio.source_code, portofolio.created_at, portofolio.updated_at, (SELECT json_agg(t) FROM (SELECT p_images.orders, p_images.images FROM p_images WHERE p_images.id_portfolio=portofolio.id) t) AS childs FROM portofolio WHERE id = $1;"
	row := dc.dbConn.QueryRow(context.Background(), query, id)
	err = row.Scan(
		&model.ID,
		&model.Title,
		&model.Descriptions,
		&model.ProjectInfo,
		&model.Languages,
		&model.Tech,
		&model.JobRole,
		&model.Databases,
		&model.Dates,
		&model.Platform,
		&model.Urls,
		&model.SourceCode,
		&model.CreatedAt,
		&model.UpdatedAt,
		&childs,
	)

	for _, v := range childs {
		sd := models.PortofolioImagesModel{
			Orders: int64(v["orders"].(float64)),
			Images: v["images"].(string),
		}
		model.Images = append(model.Images, sd)
	}

	responses = *model

	if err != nil {
		return
	}

	return
}
