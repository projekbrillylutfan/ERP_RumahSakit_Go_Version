package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type UserRepository interface {
	CreateUserRepository(ctx context.Context, user *entity.User) *entity.User
}