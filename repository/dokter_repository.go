package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type DokterRepository interface {
	CreateDokterRepository(ctx context.Context, dokter *entity.Dokter) *entity.Dokter
	FindByEmailDokterRepository(ctx context.Context, email string) (*entity.Dokter, error)
	FindAllDokterRepository(ctx context.Context) []*entity.Dokter
	FindByIdDokterRepository(ctx context.Context, id int64) (*entity.Dokter, error)
	UpdateDokterRepository(ctx context.Context, dokter *entity.Dokter) *entity.Dokter
	DeleteDokterRepository(ctx context.Context, dokter *entity.Dokter)
	AuthDokterRepository(ctx context.Context, email string) (*entity.Dokter, error)
}