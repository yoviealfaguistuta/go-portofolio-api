package http

import (
	"portfolio-api/configs"
	controllers "portfolio-api/internal/controllers"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PortofolioHandler struct {
	controller *controllers.PortofolioControllers
	validate   *validator.Validate
}

func NewPortofolioHandler(rApi fiber.Router, validator *validator.Validate, controller *controllers.PortofolioControllers) {
	handler := &PortofolioHandler{
		controller: controller,
		validate:   validator,
	}

	r := rApi.Group("/portofolio")
	r.Get("/", handler.List)
	r.Get("/:id", handler.Detail)
}

// List func to get list of all portofolio
// @Summary      Endpoint for Portoflio Pages (http://yoviealfaguistuta.site)
// @Description  Showing list of all portofolio data.
// @ID           portofolio-list
// @Tags         Portofolio
// @Produce      json
// @param        limit  		path      number    false   "Amount of data"
// @param        page   		path      number    false   "Paging"
// @success      200  {array}   models.PortofolioList  		"Success"
// @Failure      400  {object}  configs.RequestError        "Bad request"
// @Failure      404  {object}  configs.RequestError        "Data not found"
// @Failure      500  {object}  configs.ServerError         "Server error"
// @Security     ApiKeyAuth
// @Router       /portofolio/list [get]
func (ph *PortofolioHandler) List(c *fiber.Ctx) error {

	var __typename string = "PortofolioList"

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	data, page_count, total, offset, err := ph.controller.List(c, page, limit)

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

// Detail func to get detail of portofolio
// @Summary      Endpoint for Portoflio Pages (http://yoviealfaguistuta.site)
// @Description  Showing list of all portofolio data.
// @ID           portofolio-detail
// @Tags         Portofolio
// @Param        id  path  number  true  "Id Portofolio"
// @Produce      json
// @success      200  {array}   models.PortofolioDetail  "Success"
// @Failure      400  {object}  configs.RequestError        "Bad request"
// @Failure      404  {object}  configs.RequestError        "Data not found"
// @Failure      500  {object}  configs.ServerError         "Server error"
// @Security     ApiKeyAuth
// @Router       /portofolio/detail/{id} [get]
func (ph *PortofolioHandler) Detail(c *fiber.Ctx) error {

	var __typename string = "PortofolioDetail"

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return configs.ErrorResponse(c, err, __typename)
	}

	data, err := ph.controller.Detail(c, id)

	if err != nil {
		return configs.ErrorResponse(c, err, __typename)
	}

	return c.JSON(configs.Response{
		TypeName: __typename,
		Data:     data,
		Status:   true,
	})
}
