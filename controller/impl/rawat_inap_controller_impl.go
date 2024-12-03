package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

func NewRawatInapControllerImpl(rawatInapService service.RawatInapService, config configuration.Config) controller.RawatInapController {
	return &RawatInapControllerImpl{
		RawatInapService: rawatInapService,
		Config:           config,
	}
}

type RawatInapControllerImpl struct {
	RawatInapService service.RawatInapService
	Config           configuration.Config
}

func (ct *RawatInapControllerImpl) GetConfig() configuration.Config {
	return ct.Config
}

func (ct *RawatInapControllerImpl) CreateRawatInapController(c *fiber.Ctx) error {
	var req *dto.RawatInapModelCreateOrUpdate

	err := c.BodyParser(&req)
	exception.PanicLogging(err)

	result := ct.RawatInapService.CreateRawatInapService(c.Context(), req)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success create rawat inap",
		Data:    result,
	})
}