package impl_repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"gorm.io/gorm"
)

func NewResepRepositoryImpl(DB *gorm.DB) repository.ResepRepository {
	return &ResepRepositoryImpl{
		DB: DB,
	}
}

type ResepRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ResepRepositoryImpl) CreateResepRepository(ctx context.Context, Resep *entity.Resep) *entity.Resep {
	err := r.DB.WithContext(ctx).Create(&Resep).Error
	exception.PanicLogging(err)
	return Resep
}