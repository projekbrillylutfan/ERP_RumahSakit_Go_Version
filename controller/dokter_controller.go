package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type DokterController interface {
	CreateDokterController(c *fiber.Ctx) error
	FindAllDokterController(c *fiber.Ctx) error
	FindByIdDokerController(c *fiber.Ctx) error
	UpdateDokterController(c *fiber.Ctx) error
	DeleteDokterController(c *fiber.Ctx) error
	AuthDokterController(c *fiber.Ctx) error
	ForgotPassDokterController(c *fiber.Ctx) error
	ResetPassDokterController(c *fiber.Ctx) error


	GetConfig() configuration.Config
}