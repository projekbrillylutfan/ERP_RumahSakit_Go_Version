package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type ObatRepository interface {
	CreateObatRepository(ctx context.Context, obat *entity.Obat) *entity.Obat
	FindAllObatRepo(ctx context.Context) []*entity.Obat
	FindByIdObatRepo(ctx context.Context, id int64) (*entity.Obat, error)
	UpdateObatRepo(ctx context.Context, obat *entity.Obat) *entity.Obat
	DeleteObatRepo(ctx context.Context, obat *entity.Obat)
}