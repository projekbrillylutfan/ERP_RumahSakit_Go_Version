package impl_repository

import (
	"context"
	"errors"

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

func (r *ObatRepositoryImpl) FindAllObatRepo(ctx context.Context) []*entity.Obat {
	var obats []*entity.Obat
	r.DB.WithContext(ctx).Find(&obats)
	return obats
}

func (r *ObatRepositoryImpl) FindByIdObatRepo(ctx context.Context, id int64) (*entity.Obat, error) {
	var obat *entity.Obat
	result := r.DB.WithContext(ctx).Where("id_obat = ?", id).First(&obat)
	if result.RowsAffected == 0 {
		return nil, errors.New("obat not found")
	}
	return obat, nil
}

func (r *ObatRepositoryImpl) UpdateObatRepo(ctx context.Context, obat *entity.Obat) *entity.Obat {
	err := r.DB.WithContext(ctx).Where("id_obat = ?", obat.IDObat).Updates(obat).Error
	exception.PanicLogging(err)
	return obat
}