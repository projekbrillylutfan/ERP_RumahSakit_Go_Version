package impl_service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func NewDokterServiceImpl(dokterRepository *repository.DokterRepository, config *configuration.Config, redisService *RedisService, mailDialer *gomail.Dialer) service.DokterService {
	return &DokterServiceImpl{
		DokterRepository: *dokterRepository,
		Config:           *config,
		RedisService:     redisService,
		MailDialer:       mailDialer,
	}
}

type DokterServiceImpl struct {
	DokterRepository repository.DokterRepository
	Config           configuration.Config
	RedisService     *RedisService
	MailDialer       *gomail.Dialer
}

func (s *DokterServiceImpl) CreateDokterService(ctx context.Context, dokter *dto.DokterCreateOrUpdaeRequest) *dto.DokterCreateOrUpdaeRequest {
	configuration.Validate(dokter)

	_, err := s.DokterRepository.FindByEmailDokterRepository(ctx, dokter.Email)
	if err != nil {
		panic(exception.ConflictError{
			Message: err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dokter.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)
	
	dokterCreate := &entity.Dokter{
		Nama: dokter.Nama,
		Email: dokter.Email,
		Password: string(hashedPassword),
		Spesialisasi: dokter.Spesialisasi,
		NomorTelepon: dokter.NomorTelepon,
	}
	s.DokterRepository.CreateDokterRepository(ctx, dokterCreate)
	dokterCreate.Password = ""
	return dokter
}