package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/middleware"
)

func RawatInapRoute(app *fiber.App, ct controller.RawatInapController) {
	rawatInapGroup := app.Group("/api/rawat-inap", middleware.AuthenticateJWT([]string{"ADMIN", "PERAWAT"}, ct.GetConfig()))

	rawatInapGroup.Post("/", ct.CreateRawatInapController)
}