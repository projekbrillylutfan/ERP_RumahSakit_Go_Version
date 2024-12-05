package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type ResepDetailController interface {
	CreateResepDetailController(c *fiber.Ctx) error
	FindAllResepDetailController(c *fiber.Ctx) error
	FindByIdResepDetailController(c *fiber.Ctx) error

	GetConfig() configuration.Config
}