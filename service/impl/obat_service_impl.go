package impl_service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

func NewObatServiceImpl(obatRepository *repository.ObatRepository) service.ObatService {
	return &ObatServiceImpl{
		ObatRepository: *obatRepository,
	}
}

type ObatServiceImpl struct {
	ObatRepository repository.ObatRepository
}

func (s *ObatServiceImpl) CreateObatService(ctx context.Context, user *dto.ObatCreateOrUpdateRequest) *dto.ObatCreateOrUpdateRequest {
	configuration.Validate(user)
	obatCreate := &entity.Obat{
		NamaObat:   user.NamaObat,
		Deskripsi:  user.Deskripsi,
		Harga:      user.Harga,
	}

	s.ObatRepository.CreateObatRepository(ctx, obatCreate)
	return user
}