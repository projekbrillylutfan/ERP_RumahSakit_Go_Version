package impl_repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"gorm.io/gorm"
)

func NewJanjiTemuRepositoryImpl(DB *gorm.DB) repository.JanjiTemuRepository {
	return &JanjiTemuRepositoryImpl{
		DB: DB,
	}
}

type JanjiTemuRepositoryImpl struct {
	DB *gorm.DB
}

func (r *JanjiTemuRepositoryImpl) CreateJanjiTemuRepo(ctx context.Context, JanjiTemu *entity.JanjiTemu) *entity.JanjiTemu {
	err := r.DB.WithContext(ctx).Create(&JanjiTemu).Error
	exception.PanicLogging(err)
	return JanjiTemu
}