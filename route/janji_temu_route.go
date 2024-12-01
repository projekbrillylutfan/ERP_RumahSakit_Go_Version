package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/middleware"
)

func JanjiTemu(app *fiber.App, ct controller.JanjiTemuController) {
	janjiTemuGroup := app.Group("/api/janji-temu", middleware.AuthenticateJWT([]string{"ADMIN", "USER", "PERAWAT"}, ct.GetConfig()))

	janjiTemuGroup.Post("/", ct.CreateJanjiTemuController)
}