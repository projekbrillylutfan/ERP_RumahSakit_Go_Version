package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type PerawatService interface {
	CreatePerawatService(ctx context.Context, perawat *dto.PerawatCreateOrUpdateRequest) *dto.PerawatCreateOrUpdateRequest
	FindAllPerawatService(ctx context.Context) []*dto.PerawatFindAllRespones
	FindByIdPerawatService(ctx context.Context, id int64) *dto.PerawatFindByIdRespones
}