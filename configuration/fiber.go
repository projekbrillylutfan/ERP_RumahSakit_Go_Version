package configuration

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}