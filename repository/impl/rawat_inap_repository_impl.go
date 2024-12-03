package impl_repository

import (
	"context"
	"errors"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"gorm.io/gorm"
)

func NewRawatInapRepositoryImpl(DB *gorm.DB) repository.RawatInapRepository {
	return &RawatInapRepositoryImpl{
		DB: DB,
	}
}

type RawatInapRepositoryImpl struct {
	DB *gorm.DB
}

func (r *RawatInapRepositoryImpl) CreateRawatInapRepository(ctx context.Context, RawatInap *entity.RawatInap) *entity.RawatInap {
	err := r.DB.WithContext(ctx).Create(&RawatInap).Error
	exception.PanicLogging(err)
	return RawatInap
}

func (r *RawatInapRepositoryImpl) FindAllRawatInapRepository(ctx context.Context) []*entity.RawatInap {
	var rawatInap []*entity.RawatInap
	err := r.DB.WithContext(ctx).Find(&rawatInap).Error
	exception.PanicLogging(err)
	return rawatInap
}

func (r *RawatInapRepositoryImpl) FindByIdRawatInapRepo(ctx context.Context, id int64) (*entity.RawatInap, error) {
	var rawatInap entity.RawatInap
	result := r.DB.WithContext(ctx).Preload("User").Preload("Kamar").Where("id_rawat_inap = ?", id).First(&rawatInap)
	if result.RowsAffected == 0 {
		return nil, errors.New("rawat inap not found")
	}
	return &rawatInap, nil
}

func (r *RawatInapRepositoryImpl) UpdateRawatInapRepo(ctx context.Context, RawatInap *entity.RawatInap) *entity.RawatInap {
	err := r.DB.WithContext(ctx).Where("id_rawat_inap = ?", RawatInap.IDRawatInap).Updates(RawatInap).Error
	exception.PanicLogging(err)
	return RawatInap
}