package impl_repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"gorm.io/gorm"
)

func NewObatRepositoryImpl(DB *gorm.DB) repository.ObatRepository {
	return &ObatRepositoryImpl{
		DB: DB,
	}
}

type ObatRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ObatRepositoryImpl) CreateObatRepository(ctx context.Context, obat *entity.Obat) *entity.Obat {
	err := r.DB.WithContext(ctx).Create(&obat).Error
	exception.PanicLogging(err)
	return obat
}