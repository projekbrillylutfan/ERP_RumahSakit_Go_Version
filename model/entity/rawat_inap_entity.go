package entity

import "time"

type RawatInap struct {
	IDRawatInap   int64     `gorm:"primaryKey;autoIncrement;column:id_rawat_inap" json:"id_rawat_inap"` // ID Rawat Inap
	IDUser        int64     `gorm:"not null;column:id_user" json:"id_user"`                             // Foreign Key ke tabel User
	IDKamar       int64     `gorm:"not null;column:id_kamar" json:"id_kamar"`                           // Foreign Key ke tabel Kamar
	TanggalMasuk  time.Time `gorm:"type:date;not null" json:"tanggal_masuk"`                            // Tanggal Masuk
	TanggalKeluar time.Time `gorm:"type:date" json:"tanggal_keluar"`                                    // Tanggal Keluar (opsional)
	CreatedAt     time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`                    // Waktu Pembuatan
	UpdatedAt     time.Time `gorm:"type:timestamp;autoUpdateTime" json:"updated_at"`                    // Waktu Pembaruan

	// Relasi ke User
	User User `gorm:"foreignKey:IDUser;references:ID" json:"user"`

	// Relasi ke Kamar
	Kamar Kamar `gorm:"foreignKey:IDKamar;references:IDKamar" json:"kamar"`
}

func (ri *RawatInap) TableName() string {
	return "rawat_inap"
}
