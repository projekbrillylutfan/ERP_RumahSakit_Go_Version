package impl_repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"gorm.io/gorm"
)

func NewResepDetailRepositoryImpl(DB *gorm.DB) repository.ResepDetailRepository {
	return &ResepDetailRepositoryImpl{
		DB: DB,
	}
}

type ResepDetailRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ResepDetailRepositoryImpl) CreateResepDetailRepository(ctx context.Context, resepDetail *entity.ResepDetail) *entity.ResepDetail {
	err := r.DB.WithContext(ctx).Create(&resepDetail).Error
	exception.PanicLogging(err)
	return resepDetail
}