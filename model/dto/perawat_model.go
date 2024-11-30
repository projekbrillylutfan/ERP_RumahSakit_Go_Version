package dto

type PerawatCreateOrUpdateRequest struct {
	Nama            string `json:"nama" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	Shift           string `json:"shift" validate:"required"`
	NomorTelepon    string `json:"nomor_telepon" validate:"required"`
}

type PerawatFindAllRespones struct {
	ID              int64  `json:"id_perawat" validate:"required"`
	Nama            string `json:"nama" validate:"required"`
	Shift           string `json:"shift" validate:"required"`
}

type PerawatFindByIdRespones struct {
	ID              int64  `json:"id_perawat" validate:"required"`
	Nama            string `json:"nama" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	Shift           string `json:"shift" validate:"required"`
	NomorTelepon    string `json:"nomor_telepon" validate:"required"`
}

type PerawatLogin struct {
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
}