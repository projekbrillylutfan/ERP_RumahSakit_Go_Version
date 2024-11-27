package impl_repository

import (
	"context"
	"errors"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"gorm.io/gorm"
)

func NewDokterRepositoryImpl(DB *gorm.DB) repository.DokterRepository {
	return &DokterRepositoryImpl{
		DB: DB,
	}
}

type DokterRepositoryImpl struct {
	*gorm.DB
}

func (r *DokterRepositoryImpl) CreateDokterRepository(ctx context.Context, dokter *entity.Dokter) *entity.Dokter {
	err := r.DB.WithContext(ctx).Create(&dokter).Error
	exception.PanicLogging(err)
	return dokter
}

func (r *DokterRepositoryImpl) FindByEmailDokterRepository(ctx context.Context, email string) (*entity.Dokter, error){
	var dokter *entity.Dokter
	result := r.DB.WithContext(ctx).Where("email = ?", email).First(&dokter)

	if result.RowsAffected != 0 {
		return nil, errors.New("dokter already exists")
	}

	return dokter, nil
}

func (r *DokterRepositoryImpl) FindAllDokterRepository(ctx context.Context) []*entity.Dokter {
	var dokters []*entity.Dokter
	r.DB.WithContext(ctx).Find(&dokters)
	return dokters
}

func (r *DokterRepositoryImpl) FindByIdDokterRepository(ctx context.Context, id int64) (*entity.Dokter, error) {
	var dokter *entity.Dokter
	result := r.DB.WithContext(ctx).Where("id_dokter = ?", id).First(&dokter)
	if result.RowsAffected == 0 {
		return nil, errors.New("dokter not found")
	}
	return dokter, nil
}

func (r *DokterRepositoryImpl) UpdateDokterRepository(ctx context.Context, dokter *entity.Dokter) *entity.Dokter {
	err := r.DB.WithContext(ctx).Where("id_dokter = ?", dokter.ID).Updates(dokter).Error
	exception.PanicLogging(err)
	return dokter
}

func (r *DokterRepositoryImpl) DeleteDokterRepository(ctx context.Context, dokter *entity.Dokter) {
	err := r.DB.WithContext(ctx).Delete(&dokter).Error
	exception.PanicLogging(err)
}