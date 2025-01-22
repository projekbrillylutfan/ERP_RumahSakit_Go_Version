package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type ResepService interface {
	CreateResepService(ctx context.Context, Resep *dto.ResepCreateOrUpdate) *dto.ResepCreateOrUpdate
}