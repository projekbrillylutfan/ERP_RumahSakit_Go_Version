package entity

import "time"

type Kamar struct {
	IDKamar      int64     `gorm:"primaryKey;autoIncrement;column:id_kamar" json:"id_kamar"` // ID Kamar
	JenisKamar   string    `gorm:"type:jenis_kamar;not null" json:"jenis_kamar"`             // Jenis Kamar (ENUM)
	TarifPerHari float64   `gorm:"type:numeric(10,2);not null" json:"tarif_per_hari"`        // Tarif Per Hari
	CreatedAt    time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`          // Waktu Pembuatan
	UpdatedAt    time.Time `gorm:"type:timestamp;autoUpdateTime" json:"updated_at"`          // Waktu Pembaruan
}

func (k *Kamar) TableName() string {
	return "kamar"
}