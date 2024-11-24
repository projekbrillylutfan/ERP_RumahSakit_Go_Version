package impl_repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"gorm.io/gorm"
)

func NewUserRepository(DB *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

type UserRepositoryImpl struct {
	*gorm.DB
}

func (r *UserRepositoryImpl)CreateUserRepository(ctx context.Context, user *entity.User) *entity.User {
	err := r.DB.WithContext(ctx).Create(&user).Error
	exception.PanicLogging(err)
	return user
}