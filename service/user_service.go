package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type UserService interface {
	CreateUserService(ctx context.Context, user *dto.UserCreateOrUpdateRequest) *dto.UserCreateOrUpdateRequest
	FindByIdUserService(ctx context.Context, id int64) *dto.UserFindByIdReponse
	FindAllUserService(ctx context.Context) []*dto.UserFindAllReponse
}