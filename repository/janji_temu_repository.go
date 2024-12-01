package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type JanjiTemuRepository interface {
	CreateJanjiTemuRepo(ctx context.Context, JanjiTemu *entity.JanjiTemu) *entity.JanjiTemu
}