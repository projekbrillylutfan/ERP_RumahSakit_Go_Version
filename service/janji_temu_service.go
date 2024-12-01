package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type JanjiTemuService interface {
	CreateJanjiTemuService(ctx context.Context, JanjiTemu *dto.JanjiTemuCreateOrUpdate) *dto.JanjiTemuCreateOrUpdate
}