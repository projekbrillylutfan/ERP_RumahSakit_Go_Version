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