package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type JanjiTemuRepository interface {
	CreateJanjiTemuRepo(ctx context.Context, JanjiTemu *entity.JanjiTemu) *entity.JanjiTemu
	FindAllJanjiTemuRepo(ctx context.Context) []*entity.JanjiTemu
	FindByIdJanjiTemuRepo(ctx context.Context, id int64) (*entity.JanjiTemu, error)
	UpdateJanjiTemuRepo(ctx context.Context, JanjiTemu *entity.JanjiTemu) *entity.JanjiTemu
	DeleteJanjiTemuRepo(ctx context.Context, JanjiTemu *entity.JanjiTemu)
}