package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/middleware"
)

func DokterRouteAdmin(app *fiber.App, ct controller.DokterController) {
	dokterGroupAdmin := app.Group("/api/admin/dokter", middleware.AuthenticateJWT([]string{"ADMIN"}, ct.GetConfig()))

	dokterGroupAdmin.Post("/", ct.CreateDokterController)
}