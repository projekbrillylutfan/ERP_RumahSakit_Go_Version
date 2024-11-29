package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type PerawatRepository interface {
	CreatePerawatRepository(ctx context.Context, perawat *entity.Perawat) *entity.Perawat
	FindPerawatByEmailAndUsernameRepository(ctx context.Context, email, username string) (*entity.Perawat, error)
	FindAllPerawatRepository(ctx context.Context) []*entity.Perawat
	FindByIdPerawatRepository(ctx context.Context, id int64) (*entity.Perawat, error)
	UpdatePerawatRepositoy(ctx context.Context, perawat *entity.Perawat) *entity.Perawat
}