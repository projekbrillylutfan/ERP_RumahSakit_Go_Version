package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/middleware"
)

func ResepDetailRoute(app *fiber.App, ct controller.ResepDetailController) {
	resepDetailGroup := app.Group("/api/resep-detail", middleware.AuthenticateJWT([]string{"ADMIN", "PERAWAT"}, ct.GetConfig()))

	resepDetailGroup.Post("/", ct.CreateResepDetailController)
	resepDetailGroup.Get("/", ct.FindAllResepDetailController)
}