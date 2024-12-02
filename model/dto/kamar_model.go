package dto

import "time"

type KamarCreateOrUpdateReq struct {
	JenisKamar   string    `json:"jenis_kamar" validate:"required,oneof=STANDARD DELUXE VIP"`             // Jenis Kamar (ENUM)
	TarifPerHari float64   `json:"tarif_per_hari" validate:"required"`        // Tarif Per Hari
}

type KamarFindAllResponse struct {
	IDKamar      int64     `json:"id_kamar" validate:"required"`
	JenisKamar   string    `json:"jenis_kamar" validate:"required"`             // Jenis Kamar (ENUM)
	TarifPerHari float64   `json:"tarif_per_hari" validate:"required"`        // Tarif Per Hari
}

type KamarFindByIdResponse struct {
	IDKamar      int64     `json:"id_kamar" validate:"required"`
	JenisKamar   string    `json:"jenis_kamar" validate:"required"`             // Jenis Kamar (ENUM)
	TarifPerHari float64   `json:"tarif_per_hari" validate:"required"`      // Tarif Per Hari
	CreatedAt    time.Time `json:"created_at" validate:"required"`          // Waktu Pembuatan
	UpdatedAt    time.Time `json:"updated_at" validate:"required"`          // Waktu Pembaruan
}