package impl_service

import (
	"context"
	"fmt"
	"time"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/utils"
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

func (s *DokterServiceImpl) FindAllDokterService(ctx context.Context) (responses []*dto.DokterFindAllResponse) {
	dokters := s.DokterRepository.FindAllDokterRepository(ctx)

	for _, dokter := range dokters {
		responses = append(responses, &dto.DokterFindAllResponse{
			ID: dokter.ID,
			Nama: dokter.Nama,
			Spesialisasi: dokter.Spesialisasi,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}

func (s *DokterServiceImpl) FindByIdDokterService(ctx context.Context, id int64) *dto.DokterFindByIdResponse {
	dokterId, err := s.DokterRepository.FindByIdDokterRepository(ctx, id)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	return &dto.DokterFindByIdResponse{
		ID: dokterId.ID,
		Nama: dokterId.Nama,
		Email: dokterId.Email,
		Spesialisasi: dokterId.Spesialisasi,
		NomorTelepon: dokterId.NomorTelepon,
	}
}

func (s *DokterServiceImpl) UpdateDokterService(ctx context.Context, dokter *dto.DokterCreateOrUpdaeRequest, ID int64) *dto.DokterCreateOrUpdaeRequest {
	configuration.Validate(dokter)

	_, err := s.DokterRepository.FindByIdDokterRepository(ctx, ID)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dokter.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)

	dokterUpdate := &entity.Dokter{
		ID: ID,
		Nama: dokter.Nama,
		Email: dokter.Email,
		Password: string(hashedPassword),
		Spesialisasi: dokter.Spesialisasi,
		NomorTelepon: dokter.NomorTelepon,
	}
	resultUpdate :=s.DokterRepository.UpdateDokterRepository(ctx, dokterUpdate)
	resultUpdate.Password = ""
	return &dto.DokterCreateOrUpdaeRequest{
		Nama: resultUpdate.Nama,
		Email: resultUpdate.Email,
		Password: resultUpdate.Password,
		Spesialisasi: resultUpdate.Spesialisasi,
		NomorTelepon: resultUpdate.NomorTelepon,
	}
}

func (s *DokterServiceImpl) DeleteDokterService(ctx context.Context, id int64) {
	result, err := s.DokterRepository.FindByIdDokterRepository(ctx, id)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}
	s.DokterRepository.DeleteDokterRepository(ctx, result)
}

func (s *DokterServiceImpl) AuthDokterService(ctx context.Context, modelLogin *dto.DokterLogin) string {
	configuration.Validate(modelLogin)

	dokters, err := s.DokterRepository.AuthDokterRepository(ctx, modelLogin.Email)
	if err != nil {
		panic(
			exception.UnauthorizedError{
				Message: err.Error(),
			},
		)
	}

	err = bcrypt.CompareHashAndPassword([]byte(dokters.Password), []byte(modelLogin.Password))

	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "usename or password is incorrect",
		})
	}

	tokenJwtResult := utils.GenerateTokenJWT(dokters.Email, dokters.Role, s.Config)

	return tokenJwtResult
}

func (s *DokterServiceImpl) ForgotPassDokterService(ctx context.Context, request *dto.ForgotPasswordDokter) error {
	configuration.Validate(request)

	dokter, err := s.DokterRepository.FindByEmailDokterRepository(ctx, request.Email)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	token := utils.GenerateToken(16)

	tokenKey := fmt.Sprintf("reset_password:%s", token)
	err = s.RedisService.Set(ctx, tokenKey, dokter.ID, time.Minute*15)
	if err != nil {
		return err
	}

	fromEmail := "ayam@example.com"
	toEmail := dokter.Email
	name := dokter.Nama
	tokenEmail := token

	err = utils.SendVerificationEmail(s.MailDialer, fromEmail, toEmail, name, tokenEmail)
	if err != nil {
		return err
	}

	return nil
}