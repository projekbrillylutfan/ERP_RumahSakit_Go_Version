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

func NewResepServiceImpl(resepRepository *repository.ResepRepository, userRepository *repository.UserRepository, dokterRepository *repository.DokterRepository, resepDetailRespository *repository.ResepDetailRepository) service.ResepService {
	return &ResepServiceImpl{
		ResepRepository: *resepRepository,
		UserRepository:  *userRepository,
		DokterRepository: *dokterRepository,
		ResepDetailRepository: *resepDetailRespository,
	}
}

type ResepServiceImpl struct {
	ResepRepository repository.ResepRepository
	UserRepository  repository.UserRepository
	DokterRepository repository.DokterRepository
	ResepDetailRepository repository.ResepDetailRepository
}

func (s *ResepServiceImpl) CreateResepService(ctx context.Context, Resep *dto.ResepCreateOrUpdate) *dto.ResepCreateOrUpdate {
	configuration.Validate(Resep)

	_, err := s.UserRepository.FindByIdUserRepo(ctx, Resep.IDUser)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	_, err = s.DokterRepository.FindByIdDokterRepository(ctx, Resep.IDDokter)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	_, err = s.ResepDetailRepository.FindByIdResepDetailRepo(ctx, Resep.IDResepDetail)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	resepCreate := &entity.Resep{
		IDUser: Resep.IDUser,
		IDDokter: Resep.IDDokter,
		IDResepDetail: Resep.IDResepDetail,
		Tanggal: Resep.Tanggal,
	}

	s.ResepRepository.CreateResepRepository(ctx, resepCreate)
	return Resep
}