package dto

type DokterCreateOrUpdaeRequest struct {
	Nama            string `json:"nama" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	Spesialisasi    string `json:"spesialisasi" validate:"required"`
	NomorTelepon    string `json:"nomor_telepon" validate:"required"`
}