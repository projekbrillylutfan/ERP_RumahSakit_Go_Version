package entity

import "time"

type Dokter struct {
	ID              int64     `gorm:"primaryKey;autoIncrement;column:id_dokter" json:"id_dokter"`   // ID Dokter
	Nama            string    `gorm:"type:varchar(255);not null" json:"nama"`                       // Nama
	Email           string    `gorm:"type:varchar(255);unique;not null" json:"email"`               // Email unik
	Password        string    `gorm:"type:varchar(255);not null" json:"password"`                   // Password
	Role            string    `gorm:"type:dokter_role;default:'DOKTER';not null" json:"role"`       // Role
	IsEmailVerified bool      `gorm:"type:boolean;default:false;not null" json:"is_email_verified"` // Email terverifikasi
	Spesialisasi    string    `gorm:"type:varchar(255)" json:"spesialisasi"`                        // Spesialisasi
	NomorTelepon    string    `gorm:"type:varchar(20)" json:"nomor_telepon"`                        // Nomor Telepon
	CreatedAt       time.Time `gorm:"type:timestamp;autoCreateTime"`   // Waktu dibuat
	UpdatedAt       time.Time `gorm:"type:timestamp;autoUpdateTime"`   // Waktu diperbarui
}

// Custom Table Name (opsional jika nama tabel tidak sama dengan nama struct)
func (d *Dokter) TableName() string {
	return "dokter"
}
