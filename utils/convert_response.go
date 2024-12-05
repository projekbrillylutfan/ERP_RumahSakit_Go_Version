package utils

import (
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/dto"
	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
)

func ConvertUserToModel(user *entity.User) *dto.UserFindAllReponse {
	return &dto.UserFindAllReponse{
		ID:       user.ID,
		Nama:     user.Nama,
		Alamat:   user.Alamat,
		Username: user.Username,
		Email:    user.Email,
		// Tambahkan field lainnya jika diperlukan
	}
}

func ConvertDokterToModel(dokter *entity.Dokter) *dto.DokterFindAllResponse {
	return &dto.DokterFindAllResponse{
		ID:           dokter.ID,
		Nama:         dokter.Nama,
		Spesialisasi: dokter.Spesialisasi,
		// Tambahkan field lainnya jika diperlukan
	}
}

func ConvertKamarToModel(kamar *entity.Kamar) *dto.KamarFindAllResponse {
	return &dto.KamarFindAllResponse{
		IDKamar:      kamar.IDKamar,
		JenisKamar:   kamar.JenisKamar,
		TarifPerHari: kamar.TarifPerHari,
	}
}

func ConvertObatToModel(obat *entity.Obat) *dto.ObatFindAllRes {
	return &dto.ObatFindAllRes{
		IDObat:     obat.IDObat,
		NamaObat:   obat.NamaObat,
		Deskripsi:  obat.Deskripsi,
		Harga:      obat.Harga,
	}
}