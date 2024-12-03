package dto

import "time"

type RawatInapModelCreateOrUpdate struct {
	IDUser        int64     `json:"id_user" validate:"required"`                             // Foreign Key ke tabel User
	IDKamar       int64     `json:"id_kamar" validate:"required"`                           // Foreign Key ke tabel Kamar
	TanggalMasuk  time.Time `json:"tanggal_masuk" validate:"required"`                            // Tanggal Masuk
	TanggalKeluar time.Time `json:"tanggal_keluar" validate:"required"`                                    // Tanggal Keluar (opsional)
}

type RawatInapFindAllResponse struct {
	IDRawatInap   int64     `json:"id_rawat_inap" validate:"required"` // ID Rawat Inap
	IDUser        int64     `json:"id_user" validate:"required"`                             // Foreign Key ke tabel User
	IDKamar       int64     `json:"id_kamar" validate:"required"`                           // Foreign Key ke tabel Kamar
	TanggalMasuk  time.Time `json:"tanggal_masuk" validate:"required"`                            // Tanggal Masuk
	TanggalKeluar time.Time `json:"tanggal_keluar" validate:"required"`                                    // Tanggal Keluar (opsional)
}