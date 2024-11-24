package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type UserController interface {
	CreateUserController(c *fiber.Ctx) error


	GetConfig() configuration.Config
}