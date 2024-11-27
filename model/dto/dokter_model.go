package dto

type DokterCreateOrUpdaeRequest struct {
	Nama            string `json:"nama" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	Spesialisasi    string `json:"spesialisasi" validate:"required"`
	NomorTelepon    string `json:"nomor_telepon" validate:"required"`
}

type DokterFindAllResponse struct {
	ID              int64  `json:"id_dokter" validate:"required"`
	Nama            string `json:"nama" validate:"required"`
	Spesialisasi    string `json:"spesialisasi" validate:"required"`
}

type DokterFindByIdResponse struct {
	ID              int64  `json:"id_dokter" validate:"required"`
	Nama            string `json:"nama" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Spesialisasi    string `json:"spesialisasi" validate:"required"`
	NomorTelepon    string `json:"nomor_telepon" validate:"required"`
}

type DokterLogin struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
}