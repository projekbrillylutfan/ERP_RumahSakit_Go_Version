package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

func NewKamarControllerImpl(kamarService service.KamarService, config configuration.Config) controller.KamarController {
	return &KamarControllerImpl{
		KamarService: kamarService,
		Config:       config,
	}
}

type KamarControllerImpl struct {
	service.KamarService
	configuration.Config
}

func (ct *KamarControllerImpl) GetConfig() configuration.Config {
	return ct.Config
}

func (ct *KamarControllerImpl) CreateKamarController(c *fiber.Ctx) error {
	var req *dto.KamarCreateOrUpdateReq
	err := c.BodyParser(&req)
	exception.PanicLogging(err)

	result := ct.KamarService.CreateKamarService(c.Context(), req)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success create kamar",
		Data:    result,
	})
}