package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type PerawatController interface {
	CreatePerawatController(c *fiber.Ctx) error

	GetConfig() configuration.Config
}