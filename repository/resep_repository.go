package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type ResepRepository interface {
	CreateResepRepository(ctx context.Context, Resep *entity.Resep) *entity.Resep
}