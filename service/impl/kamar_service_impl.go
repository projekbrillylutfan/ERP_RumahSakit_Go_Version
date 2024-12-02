package impl_service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
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

func (s *KamarServiceImpl) FindAllKamarService(ctx context.Context) (responses []*dto.KamarFindAllResponse) {
	kamars := s.KamarRepository.FindAllKamarRepository(ctx)

	for _, kamar := range kamars {
		responses = append(responses, &dto.KamarFindAllResponse{
			IDKamar:      kamar.IDKamar,
			JenisKamar:   kamar.JenisKamar,
			TarifPerHari: kamar.TarifPerHari,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}

func (s *KamarServiceImpl) FindByIdKamarService(ctx context.Context, id int64) *dto.KamarFindByIdResponse {
	kamar, err := s.KamarRepository.FindByIdKamarRepository(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	return &dto.KamarFindByIdResponse{
		IDKamar:      kamar.IDKamar,
		JenisKamar:   kamar.JenisKamar,
		TarifPerHari: kamar.TarifPerHari,
	}
}