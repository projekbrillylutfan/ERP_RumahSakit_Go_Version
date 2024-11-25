package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
)

func UserRouteAdmin(app *fiber.App, ct controller.UserController) {
	userAdminGroup := app.Group("/api/admin/user")

	userAdminGroup.Post("/", ct.CreateUserController)
	userAdminGroup.Get("/", ct.FindAllUserController)
	userAdminGroup.Get("/:id", ct.FindByIdUserController)
	userAdminGroup.Put("/:id", ct.UpdateUserController)
}