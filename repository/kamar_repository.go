package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type KamarRepository interface {
	CreateKamarRepository(ctx context.Context, kamar *entity.Kamar) *entity.Kamar
	FindAllKamarRepository(ctx context.Context) []*entity.Kamar
	FindByIdKamarRepository(ctx context.Context, id int64) (*entity.Kamar, error)
	UpdateKamarRepository(ctx context.Context, kamar *entity.Kamar) *entity.Kamar
	DeleteKamarRepository(ctx context.Context, kamar *entity.Kamar)
}