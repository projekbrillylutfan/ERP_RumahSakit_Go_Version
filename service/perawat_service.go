package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type PerawatService interface {
	CreatePerawatService(ctx context.Context, perawat *dto.PerawatCreateOrUpdateRequest) *dto.PerawatCreateOrUpdateRequest
}