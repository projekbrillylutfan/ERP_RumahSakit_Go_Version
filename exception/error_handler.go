package exception

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		return ctx.Status(fiber.StatusBadRequest).JSON(&dto.GeneralResponseError{
			Code:    400,
			Message: "Bad Request",
			Error:   messages,
		})
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(&dto.GeneralResponseError{
			Code:    404,
			Message: "Not Found",
			Error:   err.Error(),
		})
	}

	_, conversionError := err.(ConversionError)
	if conversionError {
		return ctx.Status(fiber.StatusBadRequest).JSON(&dto.GeneralResponseError{
			Code:    400,
			Message: "Bad Request",
			Error:   "error convert string to int",
		})
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&dto.GeneralResponseError{
			Code:    401,
			Message: "Unauthorized",
			Error:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(&dto.GeneralResponseError{
		Code:    500,
		Message: "Internal Server Error",
		Error:   err.Error(),
	})
}
