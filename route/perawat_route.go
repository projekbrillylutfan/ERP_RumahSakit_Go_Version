package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/middleware"
)

func PerawatRouteAdmin(app *fiber.App, ct controller.PerawatController) {
	perawatGroupAdmin := app.Group("/api/admin/perawat", middleware.AuthenticateJWT([]string{"ADMIN"}, ct.GetConfig()))

	perawatGroupAdmin.Post("/", ct.CreatePerawatController)
	perawatGroupAdmin.Get("/:id", ct.FindByIdPerawatController)
}

func PerawatRoute(app *fiber.App, ct controller.PerawatController) {
	perawatGroup := app.Group("/api/perawat")

	perawatGroup.Get("/", ct.FindAllPerawatController)
}