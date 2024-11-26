package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

type DokterControllerImpl struct {
	service.DokterService
	configuration.Config
}

func NewDokterControllerImpl(dokterService service.DokterService, config configuration.Config) controller.DokterController {
	return &DokterControllerImpl{
		DokterService: dokterService,
		Config:        config,
	}
}

func (ct *DokterControllerImpl) GetConfig() configuration.Config {
	return ct.Config
}

func (ct *DokterControllerImpl) CreateDokterController(c *fiber.Ctx) error {
	var request *dto.DokterCreateOrUpdaeRequest
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	ct.DokterService.CreateDokterService(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success create dokter",
		Data:    request,
	})
}

func (ct *DokterControllerImpl) FindAllDokterController(c *fiber.Ctx) error {
	result := ct.DokterService.FindAllDokterService(c.Context())

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find all dokter",
		Data:    result,
	})
}