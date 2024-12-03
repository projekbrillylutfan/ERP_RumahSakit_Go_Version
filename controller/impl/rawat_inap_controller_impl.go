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

func (ct *RawatInapControllerImpl) FindAllRawatInapController(c *fiber.Ctx) error {
	result := ct.RawatInapService.FindAllRawatInapService(c.Context())
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find all rawat inap",
		Data:    result,
	})
}

func (ct *RawatInapControllerImpl) FindByIdRawatInapController(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt64 := exception.ConversionErrorStrconv(id)
	result := ct.RawatInapService.FindByIdRawatInapService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find rawat inap by id",
		Data:    result,
	})
}

func (ct *RawatInapControllerImpl) UpdateRawatInapController(c *fiber.Ctx) error {
	var req *dto.RawatInapModelCreateOrUpdate

	id := c.Params("id")

	err := c.BodyParser(&req)

	exception.PanicLogging(err)

	idInt64 := exception.ConversionErrorStrconv(id)

	result := ct.RawatInapService.UpdateRawatInapService(c.Context(), req, idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success update rawat inap",
		Data:    result,
	})
}