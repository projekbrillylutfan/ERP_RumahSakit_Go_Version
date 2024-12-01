package service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
)

type JanjiTemuService interface {
	CreateJanjiTemuService(ctx context.Context, JanjiTemu *dto.JanjiTemuCreateOrUpdate) *dto.JanjiTemuCreateOrUpdate
	FindAllJanjiTemuService(ctx context.Context) []*dto.JanjiTemuFindAllResponse
	FindByIdJanjiTemuService(ctx context.Context, id int64) *dto.JanjiTemuFindByIdResponse
	UpdateJanjiTemuService(ctx context.Context, JanjiTemu *dto.JanjiTemuCreateOrUpdate, id int64) *dto.JanjiTemuCreateOrUpdate
	DeleteJanjiTemuService(ctx context.Context, id int64)
}