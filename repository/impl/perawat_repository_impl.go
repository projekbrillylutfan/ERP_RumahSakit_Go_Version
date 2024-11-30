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

func (r *PerawatRepositoryImpl) FindAllPerawatRepository(ctx context.Context) []*entity.Perawat {
	var perawats []*entity.Perawat
	r.DB.WithContext(ctx).Find(&perawats)
	return perawats
}

func (r *PerawatRepositoryImpl) FindByIdPerawatRepository(ctx context.Context, id int64) (*entity.Perawat, error) {
	var perawat *entity.Perawat
	result := r.DB.WithContext(ctx).Where("id_perawat = ?", id).First(&perawat)
	if result.RowsAffected == 0 {
		return nil, errors.New("perawat not found")
	}
	return perawat, nil
}

func (r *PerawatRepositoryImpl) UpdatePerawatRepositoy(ctx context.Context, perawat *entity.Perawat) *entity.Perawat {
	err := r.DB.WithContext(ctx).Where("id_perawat = ?", perawat.ID).Updates(perawat).Error
	exception.PanicLogging(err)
	return perawat
}

func (r *PerawatRepositoryImpl) DeletePerawatRepository(ctx context.Context, perawat *entity.Perawat) {
	err := r.DB.WithContext(ctx).Delete(&perawat).Error
	exception.PanicLogging(err)
}

func (r *PerawatRepositoryImpl) FindByEmailPerawatRepository(ctx context.Context, email string) (*entity.Perawat, error) {
	var perawat *entity.Perawat
	result := r.DB.WithContext(ctx).Where("email = ?", email).First(&perawat)
	if result.RowsAffected == 0 {
		return nil, errors.New("perawat not found")
	}
	return perawat, nil
}

func (r *PerawatRepositoryImpl) UpdatePasswordPerawatRepository(ctx context.Context, userID int64, hashedPassword string) error {
	result := r.DB.WithContext(ctx).Model(&entity.Perawat{}).Where("id_perawat = ?", userID).Update("password", hashedPassword)
	return result.Error
}