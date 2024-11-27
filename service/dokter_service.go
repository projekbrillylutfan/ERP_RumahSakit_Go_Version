package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type DokterService interface {
	CreateDokterService(ctx context.Context, dokter *dto.DokterCreateOrUpdaeRequest) *dto.DokterCreateOrUpdaeRequest
	FindAllDokterService(ctx context.Context) []*dto.DokterFindAllResponse
	FindByIdDokterService(ctx context.Context, id int64) *dto.DokterFindByIdResponse
	UpdateDokterService(ctx context.Context, dokter *dto.DokterCreateOrUpdaeRequest, ID int64) *dto.DokterCreateOrUpdaeRequest
	DeleteDokterService(ctx context.Context, id int64)
	AuthDokterService(ctx context.Context, modelLogin *dto.DokterLogin) string
}