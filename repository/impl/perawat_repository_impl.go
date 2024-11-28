package impl_repository

import (
	"context"
	"errors"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"gorm.io/gorm"
)

func NewPerawatRepositoryImpl(DB *gorm.DB) repository.PerawatRepository {
	return &PerawatRepositoryImpl{
		DB: DB,
	}
}

type PerawatRepositoryImpl struct {
	DB *gorm.DB
}

func (r *PerawatRepositoryImpl) CreatePerawatRepository(ctx context.Context, perawat *entity.Perawat) *entity.Perawat {
	err := r.DB.WithContext(ctx).Create(&perawat).Error
	exception.PanicLogging(err)
	return perawat
}


func (r *PerawatRepositoryImpl) FindPerawatByEmailAndUsernameRepository(ctx context.Context, email, username string) (*entity.Perawat, error) {
	var perawat *entity.Perawat
	result := r.DB.WithContext(ctx).Where("email = ? OR username = ?", email, username).First(&perawat)

	if result.RowsAffected != 0 {
		return nil, errors.New("perawat already exists")
	}

	return perawat, nil
}