package repository

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type KamarRepository interface {
	CreateKamarRepository(ctx context.Context, kamar *entity.Kamar) *entity.Kamar
}