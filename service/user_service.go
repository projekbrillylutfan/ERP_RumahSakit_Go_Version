package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type UserService interface {
	CreateUserService(ctx context.Context, user *dto.UserCreateOrUpdateRequest) *dto.UserCreateOrUpdateRequest
	FindByIdUserService(ctx context.Context, id int64) *dto.UserFindByIdReponse
	FindAllUserService(ctx context.Context) []*dto.UserFindAllReponse
	UpdateUserService(ctx context.Context, user *dto.UserCreateOrUpdateRequest, id int64) *dto.UserCreateOrUpdateRequest
	DeleteUserService(ctx context.Context, id int64)
	RegisterUserService(ctx context.Context, user *dto.UserCreateOrUpdateRequestRegister) *dto.UserCreateOrUpdateRequestRegister
	VerifyEmailService(ctx context.Context, token string) error
}