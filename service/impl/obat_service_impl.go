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

func NewObatServiceImpl(obatRepository *repository.ObatRepository) service.ObatService {
	return &ObatServiceImpl{
		ObatRepository: *obatRepository,
	}
}

type ObatServiceImpl struct {
	ObatRepository repository.ObatRepository
}

func (s *ObatServiceImpl) CreateObatService(ctx context.Context, obat *dto.ObatCreateOrUpdateRequest) *dto.ObatCreateOrUpdateRequest {
	configuration.Validate(obat)
	obatCreate := &entity.Obat{
		NamaObat:   obat.NamaObat,
		Deskripsi:  obat.Deskripsi,
		Harga:      obat.Harga,
	}

	s.ObatRepository.CreateObatRepository(ctx, obatCreate)
	return obat
}

func (s *ObatServiceImpl) FindAllObatService(ctx context.Context) (responses []*dto.ObatFindAllRes) {
	obats := s.ObatRepository.FindAllObatRepo(ctx)

	for _, obat := range obats {
		responses = append(responses, &dto.ObatFindAllRes{
			IDObat:     obat.IDObat,
			NamaObat:   obat.NamaObat,
			Deskripsi:  obat.Deskripsi,
			Harga:      obat.Harga,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}

func (s *ObatServiceImpl) FindByIdObatService(ctx context.Context, id int64) *dto.ObatFindByIdRes {
	obat, err := s.ObatRepository.FindByIdObatRepo(ctx, id)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	return &dto.ObatFindByIdRes{
		IDObat:     obat.IDObat,
		NamaObat:   obat.NamaObat,
		Deskripsi:  obat.Deskripsi,
		Harga:      obat.Harga,
		CreatedAt:  obat.CreatedAt,
		UpdatedAt:  obat.UpdatedAt,
	}
}

func (s *ObatServiceImpl) UpdateObatService(ctx context.Context, obat *dto.ObatCreateOrUpdateRequest, id int64) *dto.ObatCreateOrUpdateRequest {
	configuration.Validate(obat)

	_, err := s.ObatRepository.FindByIdObatRepo(ctx, id)
	if err != nil {
		panic(
			exception.NotFoundError{
				Message: err.Error(),
			},
		)
	}

	obatUpdate := &entity.Obat{
		IDObat:     id,
		NamaObat:   obat.NamaObat,
		Deskripsi:  obat.Deskripsi,
		Harga:      obat.Harga,
	}
	s.ObatRepository.UpdateObatRepo(ctx, obatUpdate)
	return obat
}