package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type PerawatService interface {
	CreatePerawatService(ctx context.Context, perawat *dto.PerawatCreateOrUpdateRequest) *dto.PerawatCreateOrUpdateRequest
	FindAllPerawatService(ctx context.Context) []*dto.PerawatFindAllRespones
	FindByIdPerawatService(ctx context.Context, id int64) *dto.PerawatFindByIdRespones
	UpdatePerawatService(ctx context.Context, perawat *dto.PerawatCreateOrUpdateRequest, id int64) *dto.PerawatCreateOrUpdateRequest
	DeletePerawatService(ctx context.Context, id int64)
	AuthPerawatService(ctx context.Context, model *dto.PerawatLogin) string
	ForgotPassPerawatService(ctx context.Context, model *dto.ForgotPassword) error
	ResetPassPerawatService(ctx context.Context, request *dto.ResetPassword) error
}