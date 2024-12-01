package dto

import (
	"time"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

type JanjiTemuCreateOrUpdate struct {
	IDUser   int64     `json:"id_user" validate:"required"`   // ID User (Foreign Key)
	IDDokter int64     `json:"id_dokter" validate:"required"` // ID Dokter (Foreign Key)
	Tanggal  time.Time `json:"tanggal" validate:"required"`   // Tanggal Janji Temu
	Waktu    string    `json:"waktu" validate:"required,max=50"`
}

type JanjiTemuFindAllResponse struct {
	IDJanjiTemu int64     `json:"id_janji_temu" validate:"required"`
	IDUser      int64     `json:"id_user" validate:"required"`   // ID User (Foreign Key)
	IDDokter    int64     `json:"id_dokter" validate:"required"` // ID Dokter (Foreign Key)
	Tanggal     time.Time `json:"tanggal" validate:"required"`   // Tanggal Janji Temu
	Waktu       string    `json:"waktu" validate:"required,max=50"`
}

type JanjiTemuFindByIdResponse struct {
	IDJanjiTemu int64                  `json:"id_janji_temu" validate:"required"`
	IDUser      *UserFindAllReponse    `json:"id_user" validate:"required"`   // ID User (Foreign Key)
	IDDokter    *DokterFindAllResponse `json:"id_dokter" validate:"required"` // ID Dokter (Foreign Key)
	Tanggal     time.Time              `json:"tanggal" validate:"required"`   // Tanggal Janji Temu
	Waktu       string                 `json:"waktu" validate:"required,max=50"`
}

func ConvertUserToModel(user *entity.User) *UserFindAllReponse {
	return &UserFindAllReponse{
		ID:       user.ID,
		Nama:     user.Nama,
		Alamat:   user.Alamat,
		Username: user.Username,
		Email:    user.Email,
		// Tambahkan field lainnya jika diperlukan
	}
}

func ConvertDokterToModel(dokter *entity.Dokter) *DokterFindAllResponse {
	return &DokterFindAllResponse{
		ID:           dokter.ID,
		Nama:         dokter.Nama,
		Spesialisasi: dokter.Spesialisasi,
		// Tambahkan field lainnya jika diperlukan
	}
}
