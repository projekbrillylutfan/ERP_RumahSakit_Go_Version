package entity

import "time"

type Perawat struct {
	ID              int64     `gorm:"primaryKey;autoIncrement;column:id_perawat" json:"id_perawat"`
	Nama            string    `gorm:"type:varchar(255);not null" json:"nama"`
	Email           string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Username        string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password        string    `gorm:"type:varchar(255);not null" json:"password"`
	Role            string    `gorm:"type:perawat_role;default:'PERAWAT';not null" json:"role"`
	Shift           string    `gorm:"type:varchar(100);not null" json:"shift"`
	NomorTelepon    string    `gorm:"type:varchar(20)" json:"nomor_telepon"`
	CreatedAt       time.Time `gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt       time.Time `gorm:"type:timestamp;autoUpdateTime"`
}

func (p *Perawat) TableName() string {
	return "perawat"
}