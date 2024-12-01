package impl_repository

import (
	"context"
	"errors"

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

func (r *JanjiTemuRepositoryImpl) FindAllJanjiTemuRepo(ctx context.Context) []*entity.JanjiTemu {
	var JanjiTemu []*entity.JanjiTemu
	r.DB.WithContext(ctx).Find(&JanjiTemu)
	return JanjiTemu
}

func (r *JanjiTemuRepositoryImpl) FindByIdJanjiTemuRepo(ctx context.Context, id int64) (*entity.JanjiTemu, error) {
	var JanjiTemu *entity.JanjiTemu
	result := r.DB.WithContext(ctx).Unscoped().Preload("User").Preload("Dokter").Where("id_janji_temu = ?", id).First(&JanjiTemu)
	if result.Error != nil {
		return nil, errors.New("janji temu not found")
	}
	return JanjiTemu, nil
}