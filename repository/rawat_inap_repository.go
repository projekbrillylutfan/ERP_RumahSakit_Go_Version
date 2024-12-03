package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type RawatInapRepository interface {
	CreateRawatInapRepository(ctx context.Context, RawatInap *entity.RawatInap) *entity.RawatInap
	FindAllRawatInapRepository(ctx context.Context) []*entity.RawatInap
	FindByIdRawatInapRepo(ctx context.Context, id int64) (*entity.RawatInap, error)
	UpdateRawatInapRepo(ctx context.Context, RawatInap *entity.RawatInap) *entity.RawatInap
	DeleteRawatInapRepo(ctx context.Context, RawatInap *entity.RawatInap)
}