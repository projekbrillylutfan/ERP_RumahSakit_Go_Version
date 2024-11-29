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

func NewPerawatServiceImpl(perawatRepository *repository.PerawatRepository, config *configuration.Config, redisService *RedisService, mailDialer *gomail.Dialer) service.PerawatService {
	return &PerawatServiceImpl{
		PerawatRepository: *perawatRepository,
		Config:            *config,
		RedisService:      redisService,
		MailDialer:        mailDialer,
	}
}

type PerawatServiceImpl struct {
	PerawatRepository repository.PerawatRepository
	Config            configuration.Config
	RedisService      *RedisService
	MailDialer        *gomail.Dialer
}

func (s *PerawatServiceImpl) CreatePerawatService(ctx context.Context, perawat *dto.PerawatCreateOrUpdateRequest) *dto.PerawatCreateOrUpdateRequest {
	configuration.Validate(perawat)
	_, err := s.PerawatRepository.FindPerawatByEmailAndUsernameRepository(ctx, perawat.Email, perawat.Username)
	if err != nil {
		panic(
			exception.ConflictError{
				Message: err.Error(),
			},
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(perawat.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)

	perawatCreate := &entity.Perawat{
		Nama: perawat.Nama,
		Email: perawat.Email,
		Username: perawat.Username,
		Password: string(hashedPassword),
		Shift: perawat.Shift,
		NomorTelepon: perawat.NomorTelepon,
	}

	s.PerawatRepository.CreatePerawatRepository(ctx, perawatCreate)
	perawatCreate.Password = ""
	return perawat
}

func (s *PerawatServiceImpl) FindAllPerawatService(ctx context.Context) (responses []*dto.PerawatFindAllRespones) {
	perawats := s.PerawatRepository.FindAllPerawatRepository(ctx)

	for _, perawat := range perawats {
		responses = append(responses, &dto.PerawatFindAllRespones{
			ID: perawat.ID,
			Nama: perawat.Nama,
			Shift: perawat.Shift,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}