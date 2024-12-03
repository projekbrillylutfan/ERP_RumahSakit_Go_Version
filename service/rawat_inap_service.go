package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type RawatInapService interface {
	CreateRawatInapService(ctx context.Context, RawatInap *dto.RawatInapModelCreateOrUpdate) *dto.RawatInapModelCreateOrUpdate
}