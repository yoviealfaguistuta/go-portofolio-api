package configs

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"
)

type PaginationResponse struct {
	Error    []Errors    `json:"error,omitempty"`
	TypeName string      `json:"__typename"`
	Data     interface{} `json:"data,omitempty"`
	Status   bool        `json:"status"`
	Meta     Pagination  `json:"meta,omitempty"`
}

type Response struct {
	Error    []Errors    `json:"error,omitempty"`
	TypeName string      `json:"__typename"`
	Data     interface{} `json:"data,omitempty"`
	Status   bool        `json:"status"`
}

type Errors struct {
	Message string `json:"message" example:"nama: Harus diisi, tidak boleh kosong"`
}

type Pagination struct {
	PageCount   int `json:"page_count" example:"1"`
	Total       int `json:"total" example:"1"`
	CurrentPage int `json:"current_page" example:"1"`
	Offset      int `json:"offset" example:"1"`
}

func (dve Errors) Error() string {
	return dve.Message
}

type ServerErrors struct {
	Message string `json:"message" example:"schema: table is not exists"`
}

func (dve ServerErrors) Error() string {
	return dve.Message
}

type SwaggerSuccessResponse struct {
	TypeName string      `json:"__typename"`
	Data     interface{} `json:"data"`
	Status   bool        `json:"status"`
}

func ErrorResponse(ctx *fiber.Ctx, err error, typename string) error {
	status := fiber.StatusInternalServerError

	if _, ok := err.(Errors); ok {
		status = fiber.StatusUnprocessableEntity

		return ctx.Status(status).JSON(Response{
			Error: []Errors{{
				Message: err.Error(),
			}},
			TypeName: typename,
			Status:   false,
		})

	} else if _, ok := err.(*pgconn.PgError); ok {

		return ctx.Status(status).JSON(Response{
			Error: []Errors{{
				Message: err.(*pgconn.PgError).Message,
			}},
			TypeName: typename,
			Status:   false,
		})

	}

	unknownError := fmt.Sprintf("%s", err)
	return ctx.Status(status).JSON(Response{
		Error: []Errors{{
			Message: unknownError,
		}},
		TypeName: typename,
		Status:   false,
	})
}
