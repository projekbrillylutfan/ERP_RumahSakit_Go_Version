package impl_service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
)

func NewKamarServiceImpl(kamarRepository *repository.KamarRepository) service.KamarService {
	return &KamarServiceImpl{
		KamarRepository: *kamarRepository,
	}
}

type KamarServiceImpl struct {
	KamarRepository repository.KamarRepository
}

func (s *KamarServiceImpl) CreateKamarService(ctx context.Context, kamar *dto.KamarCreateOrUpdateReq) *dto.KamarCreateOrUpdateReq {
	configuration.Validate(kamar)

	kamarCreate := &entity.Kamar{
		JenisKamar:   kamar.JenisKamar,
		TarifPerHari: kamar.TarifPerHari,
	}
	s.KamarRepository.CreateKamarRepository(ctx, kamarCreate)

	return kamar
}