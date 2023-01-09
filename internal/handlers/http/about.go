package http

import (
	"portfolio-api/configs"
	controllers "portfolio-api/internal/controllers"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AboutHandler struct {
	controller *controllers.AboutControllers
	validate   *validator.Validate
}

func NewAboutHandler(rApi fiber.Router, validator *validator.Validate, controller *controllers.AboutControllers) {
	handler := &AboutHandler{
		controller: controller,
		validate:   validator,
	}

	r := rApi.Group("/about")
	r.Get("/", handler.About)
}

// About func to get detail of about tables
// @Summary      Endpoint for Portoflio Pages (http://yoviealfaguistuta.site)
// @Description  Showing detail bio information .
// @ID           about
// @Tags         About
// @Produce      json
// @success      200  {array}   models.About  "Success"
// @Failure      400  {object}  configs.RequestError        "Bad request"
// @Failure      404  {object}  configs.RequestError        "Data not found"
// @Failure      500  {object}  configs.ServerError         "Server error"
// @Security     ApiKeyAuth
// @Router       /about [get]
func (ka *AboutHandler) About(c *fiber.Ctx) error {

	var __typename string = "About"

	data, err := ka.controller.About(c)

	if err != nil {
		return configs.ErrorResponse(c, err, __typename)
	}

	return c.JSON(configs.Response{
		TypeName: __typename,
		Data:     data,
		Status:   true,
	})
}
