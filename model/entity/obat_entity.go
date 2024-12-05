package entity

import "time"

type Obat struct {
	IDObat     int64     `gorm:"primaryKey;autoIncrement;column:id_obat" json:"id_obat"` // ID Obat
	NamaObat   string    `gorm:"type:varchar(255);not null" json:"nama_obat"`            // Nama Obat
	Deskripsi  string    `gorm:"type:text" json:"deskripsi"`                            // Deskripsi Obat
	Harga      float64   `gorm:"type:numeric(10,2);not null" json:"harga"`              // Harga Obat
	CreatedAt  time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`       // Waktu Pembuatan
	UpdatedAt  time.Time `gorm:"type:timestamp;autoUpdateTime" json:"updated_at"`       // Waktu Pembaruan

	ResepDetails []*ResepDetail `gorm:"foreignKey:IDObat;references:IDObat" json:"resep_details"`
}

func (o *Obat) TableName() string {
	return "obat"
}