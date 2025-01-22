package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

func NewResepControllerImpl(resepService service.ResepService, config configuration.Config) controller.ResepController {
	return &ResepControllerImpl {
		ResepService : resepService,
		Config : config,
	}
}

type ResepControllerImpl struct {
	ResepService service.ResepService
	Config configuration.Config
}

func (ct *ResepControllerImpl) GetConfig() configuration.Config {
	return ct.Config
}

func (ct *ResepControllerImpl) CreateResepController(c *fiber.Ctx) error {
	var req *dto.ResepCreateOrUpdate
	err := c.BodyParser(&req)
	exception.PanicLogging(err)

	ct.ResepService.CreateResepService(c.Context(), req)
	return c.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code: 200,
		Message: "success create resep",
		Data: req,
	})
}