package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type DokterController interface {
	CreateDokterController(c *fiber.Ctx) error


	GetConfig() configuration.Config
}