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

func NewUserServiceImpl(userRepository *repository.UserRepository, config *configuration.Config, redisService *RedisService, mailDialer *gomail.Dialer) service.UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
		Config:         *config,
		RedisService:   redisService,
		MailDialer:     mailDialer,
	}
}

type UserServiceImpl struct {
	repository.UserRepository
	configuration.Config
	*RedisService
	MailDialer *gomail.Dialer
}

func (s *UserServiceImpl)CreateUserService(ctx context.Context, user *dto.UserCreateOrUpdateRequest) *dto.UserCreateOrUpdateRequest {
	configuration.Validate(user)
	_, err := s.UserRepository.FindByUsernamePhoneAndEmail(ctx, user.Username, user.NomorTelepon, user.Email)
	if err != nil {
		panic(exception.ConflictError{
			Message: err.Error(),
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)

	userCreate := &entity.User{
		Nama: user.Nama,
		Alamat: user.Alamat,
		Username: user.Username,
		Email: user.Email,
		Password: string(hashedPassword),
		Role: user.Role,
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		NomorTelepon: user.NomorTelepon,
	}
	s.UserRepository.CreateUserRepository(ctx, userCreate)
	userCreate.Password = ""
	return user
}

func (s *UserServiceImpl) FindByIdUserService(ctx context.Context, id int64) *dto.UserFindByIdReponse {
	userId, err := s.UserRepository.FindByIdUserRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return &dto.UserFindByIdReponse{
		ID: userId.ID,
		Nama: userId.Nama,
		Alamat: userId.Alamat,
		Username: userId.Username,
		Email: userId.Email,
		TanggalLahir: userId.TanggalLahir,
		JenisKelamin: userId.JenisKelamin,
		NomorTelepon: userId.NomorTelepon,
	}
}

func (s *UserServiceImpl) FindAllUserService(ctx context.Context) (responses []*dto.UserFindAllReponse) {
	users := s.UserRepository.FindAllUserRepo(ctx)

	for _, user := range users {
		responses = append(responses, &dto.UserFindAllReponse{
			ID: user.ID,
			Nama: user.Nama,
			Alamat: user.Alamat,
			Username: user.Username,
			Email: user.Email,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}

func (s *UserServiceImpl) UpdateUserService(ctx context.Context, user *dto.UserCreateOrUpdateRequest, id int64) *dto.UserCreateOrUpdateRequest {
	configuration.Validate(user)

	_, err := s.UserRepository.FindByIdUserRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	_, err = s.UserRepository.FindByUsernamePhoneAndEmail(ctx, user.Username, user.NomorTelepon, user.Email)
	if err != nil {
		panic(exception.ConflictError{
			Message: err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)

	userUpdate := &entity.User{
		ID: id,
		Nama: user.Nama,
		Alamat: user.Alamat,
		Username: user.Username,
		Email: user.Email,
		Password: string(hashedPassword),
		Role: user.Role,
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		NomorTelepon: user.NomorTelepon,
	}

	result := s.UserRepository.UpdateUserRepo(ctx, userUpdate)

	return &dto.UserCreateOrUpdateRequest{
		Nama: result.Nama,
		Alamat: result.Alamat,
		Username: result.Username,
		Email: result.Email,
		Password: result.Password,
		Role: result.Role,
		TanggalLahir: result.TanggalLahir,
		JenisKelamin: result.JenisKelamin,
		NomorTelepon: result.NomorTelepon,
	}
}