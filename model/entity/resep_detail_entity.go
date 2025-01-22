package entity

import "time"

type ResepDetail struct {
	IDResepDetail int64     `gorm:"primaryKey;autoIncrement;column:id_resep_detail" json:"id_resep_detail"` // ID Resep Detail
	IDObat        int64     `gorm:"not null;column:id_obat" json:"id_obat"`                                 // Foreign Key ke tabel Obat
	Jumlah        int       `gorm:"type:int;not null" json:"jumlah"`                                        // Jumlah
	Harga         float64   `gorm:"type:numeric(10,2);not null" json:"harga"`                               // Harga
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`                                       // Waktu Pembuatan
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`                                       // Waktu Pembaruan

	// Relasi ke Obat
	Obat  *Obat    `gorm:"foreignKey:IDObat;references:IDObat" json:"obat"`
	Resep []*Resep `gorm:"foreignKey:IDResepDetail;references:IDResepDetail"`
}

func (RD *ResepDetail) TableName() string {
	return "resep_detail"
}
