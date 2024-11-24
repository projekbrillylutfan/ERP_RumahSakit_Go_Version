package entity

import "time"

type User struct {
	ID              int64     `gorm:"primaryKey;autoIncrement;column:id_user"` // ID User dengan auto increment
	Nama            string    `gorm:"type:varchar(255);not null"`              // Nama
	Alamat          string    `gorm:"type:text"`                               // Alamat
	Username        string    `gorm:"type:varchar(255);unique;not null"`       // Username unik
	Email           string    `gorm:"type:varchar(255);unique;not null"`       // Email unik
	Password        string    `gorm:"type:varchar(255);not null"`              // Password
	Role            string    `gorm:"type:user_role;default:'USER';not null"`  // Role dengan default 'USER'
	IsEmailVerified bool      `gorm:"type:boolean;default:false;not null"`     // Email terverifikasi
	TanggalLahir    time.Time `gorm:"type:date"`                               // Tanggal Lahir
	JenisKelamin    string    `gorm:"type:varchar(50)"`                        // Jenis Kelamin
	NomorTelepon    string    `gorm:"type:varchar(20);unique"`                 // Nomor Telepon unik
	CreatedAt       time.Time `gorm:"type:timestamp;autoCreateTime"`           // Waktu dibuat
	UpdatedAt       time.Time `gorm:"type:timestamp;autoUpdateTime"`           // Waktu diperbarui
}

func (u *User) TableName() string {
	return "user"
}
