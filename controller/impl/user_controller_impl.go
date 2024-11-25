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

func (controller *UserControllerImpl) GetConfig() configuration.Config {
	return controller.Config
}

func (ct *UserControllerImpl) CreateUserController(c *fiber.Ctx) error {
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

func (ct *UserControllerImpl) FindByIdUserController(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt64 := exception.ConversionErrorStrconv(id)
	result := ct.UserService.FindByIdUserService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find user by id",
		Data:    result,
	})
}

func (ct *UserControllerImpl) FindAllUserController(c *fiber.Ctx) error {
	result := ct.UserService.FindAllUserService(c.Context())
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find all user",
		Data:    result,
	})
}

func (ct *UserControllerImpl) UpdateUserController(c *fiber.Ctx) error {
	var request *dto.UserCreateOrUpdateRequest
	id := c.Params("id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	idInt64 := exception.ConversionErrorStrconv(id)

	ct.UserService.UpdateUserService(c.Context(), request, idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success update user",
		Data:    request,
	})
}

func (ct *UserControllerImpl) DeleteUserController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)

	ct.UserService.DeleteUserService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success delete user",
	})
}

func (ct *UserControllerImpl) RegisterUserController(c *fiber.Ctx) error {
	var request *dto.UserCreateOrUpdateRequestRegister
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := ct.UserService.RegisterUserService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success register user, please check your email for verification",
		Data:    result,
	})
}

func (ct *UserControllerImpl) VerifyEmail(c *fiber.Ctx) error {
	token := c.Query("token")
	if token == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Token is required")
	}

	err := ct.UserService.VerifyEmailService(c.Context(), token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success verify email",
	})
}

func (ct *UserControllerImpl) LoginUserController(c *fiber.Ctx) error {
	var request *dto.UserLogin
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := ct.UserService.Authentication(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code: 200,
		Message: "success login user",
		Data: result,
	})
}
