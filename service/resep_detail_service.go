package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type ResepDetailService interface {
	CreateResepDetailService(ctx context.Context, ResepDetail *dto.ResepDetailCreateOrUpdate) *dto.ResepDetailCreateOrUpdate
}