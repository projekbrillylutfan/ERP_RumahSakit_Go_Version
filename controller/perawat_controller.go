package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type PerawatController interface {
	CreatePerawatController(c *fiber.Ctx) error
	FindAllPerawatController(c *fiber.Ctx) error
	FindByIdPerawatController(c *fiber.Ctx) error
	UpdatePerawatController(c *fiber.Ctx) error
	DeletePerawatController(c *fiber.Ctx) error
	AuthPerawatController(c *fiber.Ctx) error
	ForgotPassPerawatController(c *fiber.Ctx) error
	ResetPassPerawatController(c *fiber.Ctx) error

	GetConfig() configuration.Config
}