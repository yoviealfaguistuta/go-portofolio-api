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
	query_list := "SELECT DISTINCT on (portfolio.id) portfolio.id, portfolio.title, portfolio.description, portfolio_images.images FROM portfolio INNER JOIN portfolio_images ON portfolio.id = portfolio_images.id_portfolio order by portfolio.id, portfolio_images.id ASC"
	var rows pgx.Rows
	rows, err = dc.dbConn.Query(context.Background(), query_list)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var model = new(models.PortofolioList)
		err = rows.Scan(
			&model.ID,
			&model.Title,
			&model.Description,
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

func (dc *PortofolioControllers) Detail(c *fiber.Ctx, id int) (responses []models.PortofolioDetail, err error) {
	var model = new(models.PortofolioDetail)
	var childs []map[string]interface{}
	query_list := "SELECT portfolio.id, portfolio.title, portfolio.description, portfolio.project_info, portfolio.languages, portfolio.database, portfolio.date, portfolio.platform, portfolio.url, portfolio.source_code, portfolio.created_at, portfolio.updated_at, (SELECT json_agg(t) FROM (SELECT portfolio_images.orders, portfolio_images.images FROM portfolio_images WHERE portfolio_images.id_portfolio=portfolio.id) t) AS childs FROM portfolio WHERE id = $1;"
	row := dc.dbConn.QueryRow(context.Background(), query_list, id)
	err = row.Scan(
		&model.ID,
		&model.Title,
		&model.Description,
		&model.Database,
		&model.Dates,
		&model.Languages,
		&model.Platform,
		&model.ProjectInfo,
		&model.SourceCode,
		&model.Url,
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

	responses = append(responses, *model)

	if err != nil {
		return
	}

	return
}
