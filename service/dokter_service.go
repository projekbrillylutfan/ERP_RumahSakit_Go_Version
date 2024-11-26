package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type DokterService interface {
	CreateDokterService(ctx context.Context, dokter *dto.DokterCreateOrUpdaeRequest) *dto.DokterCreateOrUpdaeRequest 
}