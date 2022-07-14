package configs

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"
)

// NewHttpError custom error response
func NewHttpError(ctx *fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError

	if _, ok := err.(DataValidationError); ok {
		status = fiber.StatusUnprocessableEntity
		var errArr = [1]DataValidationError{
			{Field: err.(DataValidationError).Field, Message: err.Error()},
		}
		return ctx.Status(status).JSON(errArr)
	} else if _, ok := err.(RequestError); ok {
		status = fiber.StatusBadRequest
		return ctx.Status(status).JSON(RequestError{
			Code:    status,
			Message: err.Error(),
		})
	} else if _, ok := err.(ResponseError); ok {
		status = err.(ResponseError).Code
		return ctx.Status(status).JSON(RequestError{
			Code:    status,
			Message: err.Error(),
		})
	} else if _, ok := err.(ServerError); ok {
		return ctx.Status(status).JSON(ServerError{
			Code:    status,
			Message: err.Error(),
		})
	} else if _, ok := err.(*pgconn.PgError); ok {
		return ctx.Status(status).JSON(fiber.Map{"code": err.(*pgconn.PgError).Code, "message": err.(*pgconn.PgError).Message})
	} else if _, ok := err.(validator.ValidationErrors); ok {
		var fields []DataValidationError
		for _, err := range err.(validator.ValidationErrors) {
			/*fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())*/

			// penyesuaian pesan error berdasarkan jenis validasinya
			var errMsg string
			switch err.Tag() {
			case "e164":
				errMsg = "Invalid phone number format (E.164)"
			case "alphanumspace":
				errMsg = "Only alphanumeric and space allowed"
			case "alphanumslashasterisk":
				errMsg = "Only alphanumeric, slash and asterisk allowed"
			default:
				errMsg = err.Tag()
			}

			fields = append(fields, DataValidationError{Field: err.Field(), Message: errMsg})
		}
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fields)
	}

	return ctx.Status(status).JSON(fiber.Map{"code": "", "message": err.Error()})
}
