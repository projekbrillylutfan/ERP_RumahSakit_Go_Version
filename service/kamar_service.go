package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type KamarService interface {
	CreateKamarService(ctx context.Context, kamar *dto.KamarCreateOrUpdateReq) *dto.KamarCreateOrUpdateReq
}