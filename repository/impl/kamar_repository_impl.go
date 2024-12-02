package impl_repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"gorm.io/gorm"
)

func NewKamarRepositoryImpl(DB *gorm.DB) repository.KamarRepository {
	return &KamarRepositoryImpl{
		DB: DB,
	}
}

type KamarRepositoryImpl struct {
	*gorm.DB
}

func (r *KamarRepositoryImpl) CreateKamarRepository(ctx context.Context, kamar *entity.Kamar) *entity.Kamar {
	err := r.DB.WithContext(ctx).Create(&kamar).Error
	exception.PanicLogging(err)
	return kamar
}

func (r *KamarRepositoryImpl) FindAllKamarRepository(ctx context.Context) []*entity.Kamar {
	var kamars []*entity.Kamar
	r.DB.WithContext(ctx).Find(&kamars)
	return kamars
}