package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/middleware"
)

func ObatRoute(app *fiber.App, ct controller.ObatController) {
	obatrouteGroup := app.Group("/api/obat", middleware.AuthenticateJWT([]string{"ADMIN"}, ct.GetConfig()))

	obatrouteGroup.Post("/", ct.CreateObatController)
	obatrouteGroup.Get("/", ct.FindAllObatController)
	obatrouteGroup.Get("/:id", ct.FindByIdObatController)
	obatrouteGroup.Put("/:id", ct.UpdateObatController)
}