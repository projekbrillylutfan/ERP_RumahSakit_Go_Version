package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type ResepDetailRepository interface {
	CreateResepDetailRepository(ctx context.Context, resepDetail *entity.ResepDetail) *entity.ResepDetail
	FindAllResepDetailRepo(ctx context.Context) []*entity.ResepDetail
	FindByIdResepDetailRepo(ctx context.Context, id int64) (*entity.ResepDetail, error)
}