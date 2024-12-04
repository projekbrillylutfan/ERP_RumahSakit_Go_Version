package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type ObatService interface {
	CreateObatService(ctx context.Context, user *dto.ObatCreateOrUpdateRequest) *dto.ObatCreateOrUpdateRequest
	FindAllObatService(ctx context.Context) (responses []*dto.ObatFindAllRes)
}