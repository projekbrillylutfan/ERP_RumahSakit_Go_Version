package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/middleware"
)

func ResepRoute(app *fiber.App, ct controller.ResepController) {
	resepGroup := app.Group("/api/resep", middleware.AuthenticateJWT([]string{"DOKTER", "ADMIN"}, ct.GetConfig()))

	resepGroup.Post("/", ct.CreateResepController)
}