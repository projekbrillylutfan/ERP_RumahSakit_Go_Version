package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/app"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	impl_controller "github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/controller/impl"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	impl_repository "github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository/impl"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/route"
	impl_service "github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service/impl"
	"gopkg.in/gomail.v2"
)

func main() {
	config := configuration.New()
	database := app.NewDatabase(config)

	// init redis 
	redisConfig := app.NewRedisConfig()
	redisClient := app.NewRedisClient(redisConfig)
	redisService := impl_service.NewRedisService(redisClient)

	// init gomail
	// Gomail Setup
	username := config.Get("MAILTRAP_USERNAME")
	password := config.Get("MAILTRAP_PASSWORD")
	mailDialer := gomail.NewDialer("smtp.mailtrap.io", 587, username, password)

	// seed admin user
	app.SeedAdminUser(database)

	// init user repo
	userRepository := impl_repository.NewUserRepository(database)
	dokterRepository := impl_repository.NewDokterRepositoryImpl(database)
	perawatRepository := impl_repository.NewPerawatRepositoryImpl(database)
	janjiTemuRepository := impl_repository.NewJanjiTemuRepositoryImpl(database)
	kamarRepository := impl_repository.NewKamarRepositoryImpl(database)
	rawatInapRepository := impl_repository.NewRawatInapRepositoryImpl(database)
	obatRepository := impl_repository.NewObatRepositoryImpl(database)

	// init user service
	userService := impl_service.NewUserServiceImpl(&userRepository, &config, redisService, mailDialer)
	dokterService := impl_service.NewDokterServiceImpl(&dokterRepository, &config, redisService, mailDialer)
	perawatService := impl_service.NewPerawatServiceImpl(&perawatRepository, &config, redisService, mailDialer)
	janjiTemuService := impl_service.NewJanjiTemuServiceImpl(&janjiTemuRepository, &userRepository, &dokterRepository)
	kamarService := impl_service.NewKamarServiceImpl(&kamarRepository)
	rawatInapService := impl_service.NewRawatInapServiceImpl(&rawatInapRepository, &userRepository, &kamarRepository)
	obatService := impl_service.NewObatServiceImpl(&obatRepository)

	// init user controller
	userController := impl_controller.NewUserControllerImpl(userService, config)
	dokterController := impl_controller.NewDokterControllerImpl(dokterService, config)
	perawatController := impl_controller.NewPerawatControllerImpl(perawatService, config)
	janjiTemuController := impl_controller.NewJanjiTemuControllerImpl(janjiTemuService, config)
	kamarController := impl_controller.NewKamarControllerImpl(kamarService, config)
	rawatInapController := impl_controller.NewRawatInapControllerImpl(rawatInapService, config)
	obatController := impl_controller.NewObatControllerImpl(obatService, config)

	// init fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	// init route user admin
	route.UserRouteAdmin(app, userController)
	// init route user
	route.UserRoute(app, userController)
	// init route dokter admin
	route.DokterRouteAdmin(app, dokterController)
	// init route dokter
	route.DokterRoute(app, dokterController)
	// init route perawat
	route.PerawatRouteAdmin(app, perawatController)
	route.PerawatRoute(app, perawatController)
	// init route janji temu
	route.JanjiTemu(app, janjiTemuController)
	// init route kamar
	route.KamarRoute(app, kamarController)
	// init route rawat inap
	route.RawatInapRoute(app, rawatInapController)
	// init route obat
	route.ObatRoute(app, obatController)

	// server run 
	err := app.Listen(config.Get("SERVER_PORT"))
	exception.PanicLogging(err)
}