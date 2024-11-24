package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

type UserControllerImpl struct {
	service.UserService
	configuration.Config
}

func NewUserControllerImpl(userService service.UserService, config configuration.Config) controller.UserController {
	return &UserControllerImpl{
		UserService: userService,
		Config:      config,
	}
}

func (controller *UserControllerImpl) GetConfig() configuration.Config{
	return controller.Config
}

func (ct *UserControllerImpl)CreateUserController(c *fiber.Ctx) error {
	var request *dto.UserCreateOrUpdateRequest
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	ct.UserService.CreateUserService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success create user",
		Data:    request,
	})
}
