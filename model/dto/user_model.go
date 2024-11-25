package dto

import "time"

type UserCreateOrUpdateRequest struct {
	Nama         string    `json:"nama" validate:"required"`          // Nama
	Alamat       string    `json:"alamat" validate:"required"`        // Alamat
	Username     string    `json:"username" validate:"required"`      // Username unik
	Email        string    `json:"email" validate:"required,email"`   // Email unik
	Password     string    `json:"password" validate:"required"`      // Password
	Role         string    `json:"role" validate:"required"`          // Role
	TanggalLahir time.Time `json:"tanggal_lahir" validate:"required"` // Tanggal Lahir dalam format YYYY-MM-DD
	JenisKelamin string    `json:"jenis_kelamin" validate:"required"` // Jenis Kelamin
	NomorTelepon string    `json:"nomor_telepon" validate:"required"` // Nomor Telepon unik
}

type UserFindAllReponse struct {
	ID       int64  `json:"id_user" validate:"required"`
	Nama     string `json:"nama" validate:"required"`     // Nama
	Alamat   string `json:"alamat" validate:"required"`   // Alamat
	Username string `json:"username" validate:"required"` // Username unik
	Email    string `json:"email" validate:"required,email"`
}

type UserFindByIdReponse struct {
	ID           int64     `json:"id_user" validate:"required"`
	Nama         string    `json:"nama" validate:"required"`     // Nama
	Alamat       string    `json:"alamat" validate:"required"`   // Alamat
	Username     string    `json:"username" validate:"required"` // Username unik
	Email        string    `json:"email" validate:"required"`
	TanggalLahir time.Time `json:"tanggal_lahir" validate:"required"` // Tanggal Lahir dalam format YYYY-MM-DD
	JenisKelamin string    `json:"jenis_kelamin" validate:"required"` // Jenis Kelamin
	NomorTelepon string    `json:"nomor_telepon" validate:"required"`
}
