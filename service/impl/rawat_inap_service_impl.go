package impl_service

import (
	"context"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/configuration"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/exception"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/repository"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/service"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/utils"
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

func (s *RawatInapServiceImpl) FindAllRawatInapService(ctx context.Context) (responses []*dto.RawatInapFindAllResponse) {
	rawatInaps := s.RawatInapRepository.FindAllRawatInapRepository(ctx)

	for _, rawatInap := range rawatInaps {
		responses = append(responses, &dto.RawatInapFindAllResponse{
			IDRawatInap:   rawatInap.IDRawatInap,
			IDUser:        rawatInap.IDUser,
			IDKamar:       rawatInap.IDKamar,
			TanggalMasuk:  rawatInap.TanggalMasuk,
			TanggalKeluar: rawatInap.TanggalKeluar,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}

func (s *RawatInapServiceImpl) FindByIdRawatInapService(ctx context.Context, id int64) *dto.RawatInapFindByIdResponse {
	rawatInap, err := s.RawatInapRepository.FindByIdRawatInapRepo(ctx, id)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	return &dto.RawatInapFindByIdResponse{
		IDRawatInap:   rawatInap.IDRawatInap,
		IDUser:        utils.ConvertUserToModel(&rawatInap.User),
		IDKamar:       utils.ConvertKamarToModel(&rawatInap.Kamar),
		TanggalMasuk:  rawatInap.TanggalMasuk,
		TanggalKeluar: rawatInap.TanggalKeluar,
	}
}

func (s *RawatInapServiceImpl) UpdateRawatInapService(ctx context.Context, RawatInap *dto.RawatInapModelCreateOrUpdate, id int64) *dto.RawatInapModelCreateOrUpdate {
	configuration.Validate(RawatInap)

	_, err :=s.RawatInapRepository.FindByIdRawatInapRepo(ctx, id)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	_, err = s.UserRepository.FindByIdUserRepo(ctx, RawatInap.IDUser)
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

	rawatInapUpdate := &entity.RawatInap{
		IDRawatInap:   id,
		IDUser:        RawatInap.IDUser,
		IDKamar:       RawatInap.IDKamar,
		TanggalMasuk:  RawatInap.TanggalMasuk,
		TanggalKeluar: RawatInap.TanggalKeluar,
	}

	s.RawatInapRepository.UpdateRawatInapRepo(ctx, rawatInapUpdate)

	return &dto.RawatInapModelCreateOrUpdate{
		IDUser:        RawatInap.IDUser,
		IDKamar:       RawatInap.IDKamar,
		TanggalMasuk:  RawatInap.TanggalMasuk,
		TanggalKeluar: RawatInap.TanggalKeluar,
	}
}

func (s *RawatInapServiceImpl) DeleteRawatInapService(ctx context.Context, id int64) {
	result, err := s.RawatInapRepository.FindByIdRawatInapRepo(ctx, id)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}
	s.RawatInapRepository.DeleteRawatInapRepo(ctx, result)
}