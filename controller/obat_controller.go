package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type ObatController interface {
	CreateObatController(c *fiber.Ctx) error
	FindAllObatController(c *fiber.Ctx) error
	FindByIdObatController(c *fiber.Ctx) error

	GetConfig() configuration.Config
}