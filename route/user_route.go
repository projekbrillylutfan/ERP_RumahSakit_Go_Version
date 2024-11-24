package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
)

func UserRoute(app *fiber.App, ct controller.UserController) {
	userGroup := app.Group("/api/user")

	userGroup.Post("/", ct.CreateUserController)
}