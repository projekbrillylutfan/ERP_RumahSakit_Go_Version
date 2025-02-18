package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

func NewPerawatControllerImpl(perawatService service.PerawatService, config configuration.Config) controller.PerawatController {
	return &PerawatControllerImpl{
		PerawatService: perawatService,
		Config:         config,
	}
}

type PerawatControllerImpl struct {
	service.PerawatService
	configuration.Config
}

func (controller *PerawatControllerImpl) GetConfig() configuration.Config {
	return controller.Config
}

func (ct *PerawatControllerImpl) CreatePerawatController(c *fiber.Ctx) error {
	var request *dto.PerawatCreateOrUpdateRequest

	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	ct.PerawatService.CreatePerawatService(c.Context(), request)

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success create perawat",
		Data:    request,
	})
}

func (ct *PerawatControllerImpl) FindAllPerawatController(c *fiber.Ctx) error {
	result := ct.PerawatService.FindAllPerawatService(c.Context())
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find all perawat",
		Data:    result,
	})
}

func (ct *PerawatControllerImpl) FindByIdPerawatController(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt64 := exception.ConversionErrorStrconv(id)

	result := ct.PerawatService.FindByIdPerawatService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find perawat by id",
		Data:    result,
	})
}

func (ct *PerawatControllerImpl) UpdatePerawatController(c *fiber.Ctx) error {
	var request *dto.PerawatCreateOrUpdateRequest
	id := c.Params("id")

	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	idInt64 := exception.ConversionErrorStrconv(id)

	ct.PerawatService.UpdatePerawatService(c.Context(), request, idInt64)

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success update perawat",
		Data:    request,
	})
}

func (ct *PerawatControllerImpl) DeletePerawatController(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt64 := exception.ConversionErrorStrconv(id)

	ct.PerawatService.DeletePerawatService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success delete perawat",
	})
}

func (ct *PerawatControllerImpl) AuthPerawatController(c *fiber.Ctx) error {
	var requet *dto.PerawatLogin

	err := c.BodyParser(&requet)
	exception.PanicLogging(err)

	result := ct.PerawatService.AuthPerawatService(c.Context(), requet)

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success login perawat",
		Data:    result,
	})
}

func (ct *PerawatControllerImpl) ForgotPassPerawatController(c *fiber.Ctx) error {
	var request *dto.ForgotPassword

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := ct.PerawatService.ForgotPassPerawatService(c.Context(), request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "Password reset link sent to your email!",
	})
}

func (ct *PerawatControllerImpl) ResetPassPerawatController(c *fiber.Ctx) error {
	var request *dto.ResetPassword

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := ct.PerawatService.ResetPassPerawatService(c.Context(), request)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success reset password",
	})
}