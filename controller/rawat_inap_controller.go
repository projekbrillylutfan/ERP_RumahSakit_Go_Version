package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type RawatInapController interface {
	CreateRawatInapController(c *fiber.Ctx) error
	FindAllRawatInapController(c *fiber.Ctx) error
	FindByIdRawatInapController(c *fiber.Ctx) error
	UpdateRawatInapController(c *fiber.Ctx) error

	GetConfig() configuration.Config
}