package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type KamarController interface {
	CreateKamarController(c *fiber.Ctx) error
	FindAllKamarController(c *fiber.Ctx) error
	FindByIdKamarController(c *fiber.Ctx) error

	GetConfig() configuration.Config
}