package impl_service

import (
	"context"
	"fmt"
	"strconv"
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

func (s *UserServiceImpl) DeleteUserService(ctx context.Context, id int64) {
	result, err := s.UserRepository.FindByIdUserRepo(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	s.UserRepository.DeleteUserRepo(ctx, result)
}

func (s *UserServiceImpl) RegisterUserService(ctx context.Context, user *dto.UserCreateOrUpdateRequestRegister) *dto.UserCreateOrUpdateRequestRegister {
	var token string
	configuration.Validate(user)

	_, err := s.UserRepository.FindByUsernamePhoneAndEmail(ctx, user.Username, user.NomorTelepon, user.Email)
	if err != nil {
		panic(exception.ConflictError{
			Message: err.Error(),
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	exception.PanicLogging(err)

	userRegister := &entity.User{
		Nama: user.Nama,
		Alamat: user.Alamat,
		Username: user.Username,
		Email: user.Email,
		Password: string(hashedPassword),
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		NomorTelepon: user.NomorTelepon,
	}

	s.UserRepository.RegisterUserRepo(ctx, userRegister)

	token = utils.GenerateToken(16)

	tokenKey := fmt.Sprintf("email_verification:%s", token)
	err = s.RedisService.Set(ctx, tokenKey, userRegister.ID, time.Minute*15)
	exception.PanicLogging(err)

	fromEmail := "ayam@example.com"
	toEmail := userRegister.Email
	name := userRegister.Nama
	tokenEmail := token

	err = utils.SendVerificationEmail(s.MailDialer, fromEmail, toEmail, name, tokenEmail)
	if err != nil {
		panic(err.Error())
	}

	userRegister.Password = ""
	return user
}

func (s *UserServiceImpl) VerifyEmailService(ctx context.Context, token string) error {
	tokenKey := fmt.Sprintf("email_verification:%s", token)

	userID, err := s.RedisService.Get(ctx, tokenKey)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "invalid or expired token",
		})
	}

	id, _ := strconv.ParseInt(userID, 10, 64)

	err = s.UserRepository.MarkUserEmailVerified(ctx, id)
	if err != nil {
		return err
	}

	s.RedisService.Del(ctx, tokenKey)
	return nil
}

func (s *UserServiceImpl) Authentication(ctx context.Context, model *dto.UserLogin) string {
	configuration.Validate(model)
	user, err := s.UserRepository.AuthenticationRepo(ctx, model.Username)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(model.Password))

	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "usename or password is incorrect",
		})
	}

	tokenJwtResult := utils.GenerateTokenJWT(user.Username, user.Role, s.Config)

	return tokenJwtResult
}

func (s *UserServiceImpl) ForgotPasswordService(ctx context.Context, request *dto.ForgotPassword) error {
	configuration.Validate(request)
	user, err := s.FindByEmail(ctx, request.Email)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	token := utils.GenerateToken(16)

	tokenKey := fmt.Sprintf("reset_password:%s", token)
	err = s.RedisService.Set(ctx, tokenKey, user.ID, time.Minute*15)
	if err != nil {
		return err
	}

	fromEmail := "ayam@example.com"
	toEmail := user.Email
	name := user.Nama
	tokenEmail := token

	err = utils.SendVerificationEmail(s.MailDialer, fromEmail, toEmail, name, tokenEmail)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserServiceImpl) ResetPasswordService(ctx context.Context, request *dto.ResetPassword) error {
	configuration.Validate(request)
	tokenKey := fmt.Sprintf("reset_password:%s", request.Token)
	userID, err := s.RedisService.Get(ctx, tokenKey)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "invalid or expired token",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.PasswordNew), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	id, _ := strconv.ParseInt(userID, 10, 64)
	err = s.UserRepository.UpdatePassword(ctx, id, string(hashedPassword))
	if err != nil {
		return err
	}

	s.RedisService.Del(ctx, tokenKey)
	return nil
}