package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
)

type JanjiTemuController interface {
	CreateJanjiTemuController(ctx *fiber.Ctx) error
	FindAllJanjiTemuController(ctx *fiber.Ctx) error



	GetConfig() configuration.Config
}