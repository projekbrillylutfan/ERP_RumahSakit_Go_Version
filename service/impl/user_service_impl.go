package impl_service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
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
	return user
}