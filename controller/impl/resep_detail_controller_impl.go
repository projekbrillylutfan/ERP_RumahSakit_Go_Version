package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

func NewResepDetailControllerImpl(resepDetailService service.ResepDetailService, config configuration.Config) controller.ResepDetailController {
	return &ResepDetailControllerImpl{
		ResepDetailService: resepDetailService,
		Config:             config,
	}
}

type ResepDetailControllerImpl struct {
	ResepDetailService service.ResepDetailService
	Config             configuration.Config
}

func (ct *ResepDetailControllerImpl) GetConfig() configuration.Config {
	return ct.Config
}

func (ct *ResepDetailControllerImpl) CreateResepDetailController(c *fiber.Ctx) error {
	var req *dto.ResepDetailCreateOrUpdate

	err := c.BodyParser(&req)
	exception.PanicLogging(err)

	ct.ResepDetailService.CreateResepDetailService(c.Context(), req)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success create resep detail",
		Data:    req,
	})
}