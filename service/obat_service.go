package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type ObatService interface {
	CreateObatService(ctx context.Context, obat *dto.ObatCreateOrUpdateRequest) *dto.ObatCreateOrUpdateRequest
	FindAllObatService(ctx context.Context) (responses []*dto.ObatFindAllRes)
	FindByIdObatService(ctx context.Context, id int64) *dto.ObatFindByIdRes
	UpdateObatService(ctx context.Context, obat *dto.ObatCreateOrUpdateRequest, id int64) *dto.ObatCreateOrUpdateRequest
}