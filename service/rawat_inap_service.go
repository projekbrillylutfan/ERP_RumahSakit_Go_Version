package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type RawatInapService interface {
	CreateRawatInapService(ctx context.Context, RawatInap *dto.RawatInapModelCreateOrUpdate) *dto.RawatInapModelCreateOrUpdate
	FindAllRawatInapService(ctx context.Context) (responses []*dto.RawatInapFindAllResponse)
	FindByIdRawatInapService(ctx context.Context, id int64) *dto.RawatInapFindByIdResponse
	UpdateRawatInapService(ctx context.Context, RawatInap *dto.RawatInapModelCreateOrUpdate, id int64) *dto.RawatInapModelCreateOrUpdate
}