package dto

import "time"

type ResepCreateOrUpdate struct {
	IDUser        int64     `json:"id_user" validate:"required"`         // Foreign Key ke tabel User
	IDDokter      int64     `json:"id_dokter" validate:"required"`       // Foreign Key ke tabel Dokter
	IDResepDetail int64     `json:"id_resep_detail" validate:"required"` // Foreign Key ke tabel Resep Detail
	Tanggal       time.Time `json:"tanggal" validate:"required"`         // Tanggal Resep
}

type ResepFindAll struct {
	IDResep       int64     `json:"id_resep" validate:"required"`        // ID Resep (Primary Key)
	IDUser        int64     `json:"id_user" validate:"required"`         // Foreign Key ke tabel User
	IDDokter      int64     `json:"id_dokter" validate:"required"`       // Foreign Key ke tabel Dokter
	IDResepDetail int64     `json:"id_resep_detail" validate:"required"` // Foreign Key ke tabel Resep Detail
	Tanggal       time.Time `json:"tanggal" validate:"required"`         // Tanggal Resep
	TotalHarga    float64   `json:"total_harga" validate:"required"`     // Total Harga
}

type ResepFindById struct {
	IDResep       int64                  `json:"id_resep" validate:"required"`        // ID Resep (Primary Key)
	IDUser        *UserFindAllReponse    `json:"id_user" validate:"required"`         // Foreign Key ke tabel User
	IDDokter      *DokterFindAllResponse `json:"id_dokter" validate:"required"`       // Foreign Key ke tabel Dokter
	IDResepDetail *ResepDetailFindAllRes `json:"id_resep_detail" validate:"required"` // Foreign Key ke tabel Resep Detail
	Tanggal       time.Time              `json:"tanggal" validate:"required"`         // Tanggal Resep
	TotalHarga    float64                `json:"total_harga" validate:"required"`     // Total Harga
	CreatedAt     time.Time              `json:"created_at" validate:"required"`      // Waktu Pembuatan
	UpdatedAt     time.Time              `json:"updated_at" validate:"required"`      // Waktu Pembaruan
}
