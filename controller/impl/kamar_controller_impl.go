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

func (ct *KamarControllerImpl) FindAllKamarController(c *fiber.Ctx) error {
	result := ct.KamarService.FindAllKamarService(c.Context())
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find all kamar",
		Data:    result,
	})
}

func (ct *KamarControllerImpl) FindByIdKamarController(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt64 := exception.ConversionErrorStrconv(id)
	result := ct.KamarService.FindByIdKamarService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find kamar by id",
		Data:    result,
	})
}

func (ct *KamarControllerImpl) UpdateKamarController(c *fiber.Ctx) error {
	var req *dto.KamarCreateOrUpdateReq

	id := c.Params("id")
	err := c.BodyParser(&req)
	exception.PanicLogging(err)

	idInt64 := exception.ConversionErrorStrconv(id)

	result := ct.KamarService.UpdateKamarService(c.Context(), req, idInt64)

	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success update kamar",
		Data:    &result,
	})
}