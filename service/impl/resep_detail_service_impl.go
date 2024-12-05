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

func (s *ResepDetailServiceImpl) FindAllResepDetailService(ctx context.Context) (responses []*dto.ResepDetailFindAllRes) {
	resepDetail := s.ResepDetailRepository.FindAllResepDetailRepo(ctx)
	for _, resepDetail := range resepDetail {
		responses = append(responses, &dto.ResepDetailFindAllRes{
			IDResepDetail: resepDetail.IDResepDetail,
			IDObat:        resepDetail.IDObat,
			Jumlah:        resepDetail.Jumlah,
			Harga:         resepDetail.Harga,
		})
	}
	return responses
}

func (s *ResepDetailServiceImpl) FindByIdResepDetailService(ctx context.Context, id int64) *dto.ResepDetailFindByIdRes {
	resepDetail, err := s.ResepDetailRepository.FindByIdResepDetailRepo(ctx, id)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}
	return &dto.ResepDetailFindByIdRes{
		IDResepDetail: resepDetail.IDResepDetail,
		IDObat: *utils.ConvertObatToModel(resepDetail.Obat),
		Jumlah: resepDetail.Jumlah,
		Harga: resepDetail.Harga,
		CreatedAt: resepDetail.CreatedAt,
		UpdatedAt: resepDetail.UpdatedAt,
	}
}