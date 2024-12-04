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

func (ct *ObatControllerImpl) FindAllObatController(c *fiber.Ctx) error {
	result := ct.ObatService.FindAllObatService(c.Context())
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find all obat",
		Data:    result,
	})
}

func (ct *ObatControllerImpl) FindByIdObatController(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt64 := exception.ConversionErrorStrconv(id)
	result := ct.ObatService.FindByIdObatService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find obat by id",
		Data:    result,
	})
}

func (ct *ObatControllerImpl) UpdateObatController(c *fiber.Ctx) error {
	var req *dto.ObatCreateOrUpdateRequest
	id := c.Params("id")

	err := c.BodyParser(&req)
	exception.PanicLogging(err)

	idInt64 := exception.ConversionErrorStrconv(id)

	ct.ObatService.UpdateObatService(c.Context(), req, idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success update obat",
		Data:    req,
	})
}

func (ct *ObatControllerImpl) DeleteObatController(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt64 := exception.ConversionErrorStrconv(id)
	ct.ObatService.DeleteObatService(c.Context(), idInt64)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success delete obat",
	})
}