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

func NewJanjiTemuServiceImpl(janjiTemuRepository *repository.JanjiTemuRepository, userRepository *repository.UserRepository, dokterRespository *repository.DokterRepository) service.JanjiTemuService {
	return &JanjiTemuServiceImpl{
		JanjiTemuRepository: *janjiTemuRepository,
		UserRepository:      *userRepository,
		DokterRespository:   *dokterRespository,
	}
}

type JanjiTemuServiceImpl struct {
	JanjiTemuRepository repository.JanjiTemuRepository
	UserRepository      repository.UserRepository
	DokterRespository   repository.DokterRepository
}

func (s *JanjiTemuServiceImpl) CreateJanjiTemuService(ctx context.Context, JanjiTemu *dto.JanjiTemuCreateOrUpdate) *dto.JanjiTemuCreateOrUpdate {
	configuration.Validate(JanjiTemu)

	_, err := s.UserRepository.FindByIdUserRepo(ctx, JanjiTemu.IDUser)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	_, err = s.DokterRespository.FindByIdDokterRepository(ctx, JanjiTemu.IDDokter)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	janjiTemuCreate := &entity.JanjiTemu{
		IDUser:      JanjiTemu.IDUser,
		IDDokter:    JanjiTemu.IDDokter,
		Tanggal:     JanjiTemu.Tanggal,
		Waktu:       JanjiTemu.Waktu,
	}

	s.JanjiTemuRepository.CreateJanjiTemuRepo(ctx, janjiTemuCreate)
	return JanjiTemu
}

func (s *JanjiTemuServiceImpl) FindAllJanjiTemuService(ctx context.Context) (responses []*dto.JanjiTemuFindAllResponse) {
	janjiTemu := s.JanjiTemuRepository.FindAllJanjiTemuRepo(ctx)
	for _, janjiTemu := range janjiTemu {
		responses = append(responses, &dto.JanjiTemuFindAllResponse{
			IDJanjiTemu: janjiTemu.IDJanjiTemu,
			IDUser:      janjiTemu.IDUser,
			IDDokter:    janjiTemu.IDDokter,
			Tanggal:     janjiTemu.Tanggal,
			Waktu:       janjiTemu.Waktu,
		})
	}

	if len(responses) == 0 {
		return nil
	}

	return responses
}
