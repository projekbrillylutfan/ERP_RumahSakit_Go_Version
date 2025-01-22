package entity

import "time"

type Resep struct {
	IDResep       int64     `gorm:"primaryKey;autoIncrement;column:id_resep" json:"id_resep"`           // ID Resep (Primary Key)
	IDUser        int64     `gorm:"not null;column:id_user" json:"id_user"`                            // Foreign Key ke tabel User
	IDDokter      int64     `gorm:"not null;column:id_dokter" json:"id_dokter"`                        // Foreign Key ke tabel Dokter
	IDResepDetail int64     `gorm:"not null;column:id_resep_detail" json:"id_resep_detail"`            // Foreign Key ke tabel Resep Detail
	Tanggal       time.Time `gorm:"type:date;not null;column:tanggal" json:"tanggal"`                  // Tanggal Resep
	TotalHarga    float64   `gorm:"type:numeric(10,2);column:total_harga" json:"total_harga"`          // Total Harga
	CreatedAt     time.Time `gorm:"type:timestamp;autoCreateTime;column:created_at" json:"created_at"` // Waktu Pembuatan
	UpdatedAt     time.Time `gorm:"type:timestamp;autoUpdateTime;column:updated_at" json:"updated_at"` // Waktu Pembaruan

	// Relasi ke User
	User *User `gorm:"foreignKey:IDUser;references:ID" json:"user"`

	// Relasi ke Dokter
	Dokter *Dokter `gorm:"foreignKey:IDDokter;references:ID" json:"dokter"`

	// Relasi ke ResepDetail
	ResepDetail []*ResepDetail `gorm:"foreignKey:IDResepDetail;references:IDResepDetail" json:"resep_detail"`
}

func (r *Resep) TableName() string {
    return "resep"
}