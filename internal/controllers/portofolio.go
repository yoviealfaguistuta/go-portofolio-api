package controllers

import (
	"context"
	"portfolio-api/configs"
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
	var capsule []models.PortofolioList
	query_list := "SELECT DISTINCT on (portfolio.id) portfolio.id, portfolio.title, portfolio.description, portfolio_images.images FROM portfolio INNER JOIN portfolio_images ON portfolio.id = portfolio_images.id_portfolio order by portfolio.id, portfolio_images.id ASC"
	var rows pgx.Rows
	rows, err = dc.dbConn.Query(context.Background(), query_list)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var model = new(models.PortofolioList)
		err = rows.Scan(&model.ID, &model.Title, &model.Description, &model.Images)

		if err != nil {
			err = configs.RequestError{
				Code:    500,
				Message: "Gagal mengambil data Portofolio.",
			}
			return
		}
		capsule = append(capsule, *model)
	}

	if rows.Err() != nil {
		return responses, rows.Err()
	}

	return
}

func (dc *PortofolioControllers) Detail(c *fiber.Ctx, id int) (responses []models.PortofolioDetail, err error) {
	var capsule []models.PortofolioDetail
	query_list := "SELECT portfolio.id, portfolio.title, portfolio.description, portfolio_images.images FROM portfolio INNER JOIN portfolio_images ON portfolio.id = portfolio_images.id_portfolio WHERE portfolio.id = $1"
	rows, err := dc.dbConn.Query(context.Background(), query_list, id)
	if err != nil {
		err = configs.RequestError{
			Code:    500,
			Message: err.Error(),
		}
		return
	}
	defer rows.Close()

	for rows.Next() {
		var model = new(models.PortofolioDetail)
		err = rows.Scan(&model.ID, &model.Title, &model.Description, &model.Images)

		if err != nil {
			err = configs.RequestError{
				Code:    500,
				Message: "Gagal mengambil data Portofolio.",
			}
			return
		}
		capsule = append(capsule, *model)
	}

	return
}
