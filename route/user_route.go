package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/middleware"
)

func UserRouteAdmin(app *fiber.App, ct controller.UserController) {
	userAdminGroup := app.Group("/api/admin/user", middleware.AuthenticateJWT([]string{"ADMIN"}, ct.GetConfig()))

	userAdminGroup.Post("/", ct.CreateUserController)
	userAdminGroup.Get("/", ct.FindAllUserController)
	userAdminGroup.Get("/:id", ct.FindByIdUserController)
	userAdminGroup.Put("/:id", ct.UpdateUserController)
	userAdminGroup.Delete("/:id", ct.DeleteUserController)
}

func UserRoute(app *fiber.App, ct controller.UserController) {
	userGroup := app.Group("/api/user")

	userGroup.Post("/register", ct.RegisterUserController)
	userGroup.Get("/verify-email", ct.VerifyEmail)
	userGroup.Post("/login", ct.LoginUserController)
}