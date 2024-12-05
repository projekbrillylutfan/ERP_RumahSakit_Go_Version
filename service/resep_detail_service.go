package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type ResepDetailService interface {
	CreateResepDetailService(ctx context.Context, ResepDetail *dto.ResepDetailCreateOrUpdate) *dto.ResepDetailCreateOrUpdate
	FindAllResepDetailService(ctx context.Context) (responses []*dto.ResepDetailFindAllRes)
	FindByIdResepDetailService(ctx context.Context, id int64) *dto.ResepDetailFindByIdRes
	UpdateResepDetailService(ctx context.Context, ResepDetail *dto.ResepDetailCreateOrUpdate, id int64) *dto.ResepDetailCreateOrUpdate
}