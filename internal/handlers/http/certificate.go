package http

import (
	"portfolio-api/configs"
	controllers "portfolio-api/internal/controllers"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CertificateHandler struct {
	controller *controllers.CertificateControllers
	validate   *validator.Validate
}

func NewCertificateHandler(rApi fiber.Router, validator *validator.Validate, controller *controllers.CertificateControllers) {
	handler := &CertificateHandler{
		controller: controller,
		validate:   validator,
	}

	r := rApi.Group("/certificate")
	r.Get("/", handler.List)
}

// Certificate func to get list of certificates tables
// @Summary      Endpoint for Portoflio Pages (http://yoviealfaguistuta.site)
// @Description  Showing list of all certificates.
// @ID           certificate
// @Tags         Certificate
// @Produce      json
// @param        limit  		path      number    false   "Amount of data"
// @param        page   		path      number    false   "Paging"
// @success      200  {array}   models.Certificate  "Success"
// @Failure      400  {object}  configs.RequestError        "Bad request"
// @Failure      404  {object}  configs.RequestError        "Data not found"
// @Failure      500  {object}  configs.ServerError         "Server error"
// @Security     ApiKeyAuth
// @Router       /certificate [get]
func (ch *CertificateHandler) List(c *fiber.Ctx) error {

	var __typename string = "CertificateList"

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	data, page_count, total, offset, err := ch.controller.List(c, page, limit)

	if err != nil {
		return configs.ErrorResponse(c, err, __typename)
	}

	pagination := configs.Pagination{
		PageCount:   configs.Float64ToInt(page_count),
		CurrentPage: page,
		Total:       total,
		Offset:      offset,
	}

	return c.JSON(configs.PaginationResponse{
		TypeName: __typename,
		Data:     data,
		Status:   true,
		Meta:     pagination,
	})
}
