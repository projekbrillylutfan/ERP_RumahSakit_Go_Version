package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

func NewObatControllerImpl(obatService service.ObatService, config configuration.Config) controller.ObatController {
	return &ObatControllerImpl{
		ObatService: obatService,
		Config:      config,
	}
}

type ObatControllerImpl struct {
	ObatService service.ObatService
	Config      configuration.Config
}

func (ct *ObatControllerImpl) GetConfig() configuration.Config {
	return ct.Config
}

func (ct *ObatControllerImpl) CreateObatController(c *fiber.Ctx) error {
	var req *dto.ObatCreateOrUpdateRequest
	err := c.BodyParser(&req)
	exception.PanicLogging(err)

	resutl := ct.ObatService.CreateObatService(c.Context(), req)

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success create obat",
		Data:    resutl,
	})
}