package dto

import "time"

type JanjiTemuCreateOrUpdate struct {
	IDUser      int64     `json:"id_user" validate:"required"`           // ID User (Foreign Key)
	IDDokter    int64     `json:"id_dokter" validate:"required"`         // ID Dokter (Foreign Key)
	Tanggal     time.Time `json:"tanggal" validate:"required"`           // Tanggal Janji Temu
	Waktu       string    `json:"waktu" validate:"required,max=50"`
}