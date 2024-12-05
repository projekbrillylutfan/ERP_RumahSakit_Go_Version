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

func NewResepDetailServiceImpl(resepDetailRepository *repository.ResepDetailRepository, obatRepository *repository.ObatRepository) service.ResepDetailService {
	return &ResepDetailServiceImpl{
		ResepDetailRepository: *resepDetailRepository,
		ObatRepository:        *obatRepository,
	}
}

type ResepDetailServiceImpl struct {
	ResepDetailRepository repository.ResepDetailRepository
	ObatRepository        repository.ObatRepository
}

func (s *ResepDetailServiceImpl) CreateResepDetailService(ctx context.Context, ResepDetail *dto.ResepDetailCreateOrUpdate) *dto.ResepDetailCreateOrUpdate {
	configuration.Validate(ResepDetail)

	_, err := s.ObatRepository.FindByIdObatRepo(ctx, ResepDetail.IDObat)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	resepDetailCreate:= &entity.ResepDetail{
		IDObat:  ResepDetail.IDObat,
		Jumlah:  ResepDetail.Jumlah,
	}
	s.ResepDetailRepository.CreateResepDetailRepository(ctx, resepDetailCreate)
	return ResepDetail
}