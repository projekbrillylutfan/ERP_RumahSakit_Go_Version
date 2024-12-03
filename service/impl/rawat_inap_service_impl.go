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

func NewRawatInapServiceImpl(rawatInapRepository *repository.RawatInapRepository, userRepository *repository.UserRepository, kamarRepository *repository.KamarRepository) service.RawatInapService {
	return &RawatInapServiceImpl{
		RawatInapRepository: *rawatInapRepository,
		UserRepository:      *userRepository,
		KamarRepository:     *kamarRepository,
	}
}

type RawatInapServiceImpl struct {
	RawatInapRepository repository.RawatInapRepository
	UserRepository      repository.UserRepository
	KamarRepository     repository.KamarRepository
}

func (s *RawatInapServiceImpl) CreateRawatInapService(ctx context.Context, RawatInap *dto.RawatInapModelCreateOrUpdate) *dto.RawatInapModelCreateOrUpdate {
	configuration.Validate(RawatInap)
	_, err := s.UserRepository.FindByIdUserRepo(ctx, RawatInap.IDUser)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	_, err = s.KamarRepository.FindByIdKamarRepository(ctx, RawatInap.IDKamar)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	rawatInapCreate := &entity.RawatInap{
		IDUser:        RawatInap.IDUser,
		IDKamar:       RawatInap.IDKamar,
		TanggalMasuk:  RawatInap.TanggalMasuk,
		TanggalKeluar: RawatInap.TanggalKeluar,
	}

	s.RawatInapRepository.CreateRawatInapRepository(ctx, rawatInapCreate)

	return RawatInap
}