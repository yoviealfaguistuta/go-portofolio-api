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
	r.Get("/list", handler.List)
	r.Get("/detail/:id", handler.Detail)
}

// List func to get list of all portofolio
// @Summary      Endpoint for Portoflio Pages (http://yoviealfaguistuta.site)
// @Description  Showing list of all portofolio data.
// @ID           portofolio-list
// @Tags         Portofolio
// @Produce      json
// @success      200  {array}   models.PortofolioList  "Success"
// @Failure      400  {object}  configs.RequestError        "Bad request"
// @Failure      404  {object}  configs.RequestError        "Data not found"
// @Failure      422  {array}   configs.DataValidationError "Data validation failed"
// @Failure      500  {object}  configs.ServerError         "Server error"
// @Security     ApiKeyAuth
// @Router       /portofolio/list [get]
func (ka *PortofolioHandler) List(c *fiber.Ctx) error {
	var r interface{}
	var err error

	r, err = ka.controller.List(c)

	if err != nil {
		return configs.NewHttpError(c, err)
	}

	return c.Status(200).JSON(fiber.Map{
		"body":       r,
		"message":    "OK",
		"__typename": "about",
		"status":     true,
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
// @Failure      422  {array}   configs.DataValidationError "Data validation failed"
// @Failure      500  {object}  configs.ServerError         "Server error"
// @Security     ApiKeyAuth
// @Router       /portofolio/detail/{id} [get]
func (ka *PortofolioHandler) Detail(c *fiber.Ctx) error {
	var r interface{}
	var err error

	var params, convert_err = strconv.Atoi(c.Params("id"))
	if convert_err != nil {
		return configs.NewHttpError(c, convert_err)
	}

	r, err = ka.controller.Detail(c, params)

	if err != nil {
		return configs.NewHttpError(c, err)
	}

	return c.Status(200).JSON(fiber.Map{
		"body":       r,
		"message":    "OK",
		"__typename": "about",
		"status":     true,
	})
}
