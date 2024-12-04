package dto

import "time"

type ObatCreateOrUpdateRequest struct {
	NamaObat   string    `json:"nama_obat" validate:"required"`            // Nama Obat
	Deskripsi  string    `json:"deskripsi" validate:"required"`            // Deskripsi Obat
	Harga      float64   `json:"harga" validate:"required"`              // Harga Obat
}

type ObatFindAllRes struct {
	IDObat     int64     `json:"id_obat" validate:"required"`
	NamaObat   string    `json:"nama_obat" validate:"required"`
	Deskripsi  string    `json:"deskripsi" validate:"required"`
	Harga      float64   `json:"harga" validate:"required"`
}

type ObatFindByIdRes struct{
	IDObat     int64     `json:"id_obat" validate:"required"`
	NamaObat   string    `json:"nama_obat" validate:"required"`
	Deskripsi  string    `json:"deskripsi" validate:"required"`
	Harga      float64   `json:"harga" validate:"required"`
	CreatedAt  time.Time `json:"created_at" validate:"required"`
	UpdatedAt  time.Time `json:"updated_at" validate:"required"`
} 