package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type UserController interface {
	CreateUserController(c *fiber.Ctx) error
	FindByIdUserController(c *fiber.Ctx) error
	FindAllUserController(c *fiber.Ctx) error
	UpdateUserController(c *fiber.Ctx) error
	DeleteUserController(c *fiber.Ctx) error
	RegisterUserController(c *fiber.Ctx) error
	VerifyEmail(ctx *fiber.Ctx) error
	LoginUserController(c *fiber.Ctx) error
	ForgotPassword(ctx *fiber.Ctx) error
	ResetPassword(ctx *fiber.Ctx) error


	GetConfig() configuration.Config
}