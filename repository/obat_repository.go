package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type ObatRepository interface {
	CreateObatRepository(ctx context.Context, obat *entity.Obat) *entity.Obat
}