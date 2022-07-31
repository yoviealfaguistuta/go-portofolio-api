package http

import (
	"portfolio-api/configs"
	controllers "portfolio-api/internal/controllers"

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
// @Description  Showing list of all certificates data.
// @ID           certificate
// @Tags         Certificate
// @Produce      json
// @success      200  {array}   models.Certificate  "Success"
// @Failure      400  {object}  configs.RequestError        "Bad request"
// @Failure      404  {object}  configs.RequestError        "Data not found"
// @Failure      422  {array}   configs.DataValidationError "Data validation failed"
// @Failure      500  {object}  configs.ServerError         "Server error"
// @Security     ApiKeyAuth
// @Router       /certificate [get]
func (ka *CertificateHandler) List(c *fiber.Ctx) error {
	var r interface{}
	var err error

	r, err = ka.controller.List(c)

	if err != nil {
		return configs.NewHttpError(c, err)
	}

	return c.Status(200).JSON(fiber.Map{
		"body":       r,
		"message":    "OK",
		"__typename": "certificate",
		"status":     true,
	})
}
