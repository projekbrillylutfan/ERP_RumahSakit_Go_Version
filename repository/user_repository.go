package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type UserRepository interface {
	CreateUserRepository(ctx context.Context, user *entity.User) *entity.User
	FindByIdUserRepo(ctx context.Context, id int64) (*entity.User, error)
	FindAllUserRepo(ctx context.Context) []*entity.User
	FindByUsernamePhoneAndEmail(ctx context.Context, username, phone, email string) (*entity.User, error)
	UpdateUserRepo(ctx context.Context, user *entity.User) *entity.User
	DeleteUserRepo(ctx context.Context, user *entity.User)
}