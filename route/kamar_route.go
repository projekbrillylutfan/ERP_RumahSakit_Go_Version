package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/middleware"
)

func KamarRoute(app *fiber.App, ct controller.KamarController) {
	kamarGroup := app.Group("/api/kamar", middleware.AuthenticateJWT([]string{"ADMIN", "PERAWAT"}, ct.GetConfig()))

	kamarGroup.Post("/", ct.CreateKamarController)
}