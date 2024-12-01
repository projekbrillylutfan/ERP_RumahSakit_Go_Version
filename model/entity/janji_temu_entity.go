package entity

import "time"

type JanjiTemu struct {
	IDJanjiTemu int64     `gorm:"primaryKey;autoIncrement;column:id_janji_temu"` // ID Janji Temu
	IDUser      int64     `gorm:"not null;column:id_user"`                      // Foreign Key ke tabel User
	IDDokter    int64     `gorm:"not null;column:id_dokter"`                    // Foreign Key ke tabel Dokter
	Tanggal     time.Time `gorm:"type:date;not null"`                           // Tanggal Janji Temu
	Waktu       string    `gorm:"type:varchar(50);not null"`                    // Waktu Janji Temu
	CreatedAt   time.Time `gorm:"type:timestamp;autoCreateTime"`                // Waktu dibuat
	UpdatedAt   time.Time `gorm:"type:timestamp;autoUpdateTime"`                // Waktu diperbarui

	// Relasi ke User (Foreign Key: id_user)
	User User `gorm:"foreignKey:IDUser;references:ID"`

	// Relasi ke Dokter (Foreign Key: id_dokter)
	Dokter Dokter `gorm:"foreignKey:IDDokter;references:ID"`
}

func (j *JanjiTemu) TableName() string {
	return "janji_temu"
}