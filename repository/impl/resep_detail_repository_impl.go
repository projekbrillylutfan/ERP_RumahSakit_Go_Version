package impl_repository

import (
	"context"
	"errors"

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

func (r *ResepDetailRepositoryImpl) FindAllResepDetailRepo(ctx context.Context) []*entity.ResepDetail {
	var resepDetail []*entity.ResepDetail
	r.DB.WithContext(ctx).Find(&resepDetail)
	return resepDetail
}

func (r *ResepDetailRepositoryImpl) FindByIdResepDetailRepo(ctx context.Context, id int64) (*entity.ResepDetail, error) {
	var resepDetail *entity.ResepDetail
	result := r.DB.WithContext(ctx).Preload("Obat").Where("id_resep_detail = ?", id).First(&resepDetail)
	if result.RowsAffected == 0 {
		return nil, errors.New("resep detail not found")
	}
	return resepDetail, nil
}

func (r *ResepDetailRepositoryImpl) UpdateResepDetailRepo(ctx context.Context, ResepDetail *entity.ResepDetail) *entity.ResepDetail {
	err := r.DB.WithContext(ctx).Where("id_resep_detail", ResepDetail.IDResepDetail).Updates(&ResepDetail).Error
	exception.PanicLogging(err)
	return ResepDetail
}

func (r *ResepDetailRepositoryImpl) DeleteResepDetailRepo(ctx context.Context, ResepDetail *entity.ResepDetail) {
	err := r.DB.WithContext(ctx).Delete(&ResepDetail).Error
	exception.PanicLogging(err)
}