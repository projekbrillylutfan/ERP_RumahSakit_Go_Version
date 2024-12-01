package impl_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

func NewJanjiTemuControllerImpl(janjiTemuService service.JanjiTemuService, config configuration.Config) controller.JanjiTemuController {
	return &JanjiTemuControllerImpl{
		JanjiTemuService: janjiTemuService,
		Config:           config,
	}
}

type JanjiTemuControllerImpl struct {
	JanjiTemuService service.JanjiTemuService
	Config           configuration.Config
}

func (ct *JanjiTemuControllerImpl) GetConfig() configuration.Config {
	return ct.Config
}

func (ct *JanjiTemuControllerImpl) CreateJanjiTemuController(ctx *fiber.Ctx) error {
	var request *dto.JanjiTemuCreateOrUpdate
	err := ctx.BodyParser(&request)
	exception.PanicLogging(err)

	ct.JanjiTemuService.CreateJanjiTemuService(ctx.Context(), request)
	return ctx.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success create janji temu",
		Data:    request,
	})
}

func (ct *JanjiTemuControllerImpl) FindAllJanjiTemuController(ctx *fiber.Ctx) error {
	result := ct.JanjiTemuService.FindAllJanjiTemuService(ctx.Context())
	return ctx.Status(fiber.StatusOK).JSON(&dto.GeneralResponse{
		Code:    200,
		Message: "success find all janji temu",
		Data:    result,
	})
}