package app

import (
	"fmt"
	"log"
	"time"

	"github.com/projekbrillylutfan/ERP_RumahSakit_Go_Version/model/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdminUser(DB *gorm.DB) {
	// Cek apakah user dengan username 'test' sudah ada
	var user *entity.User
	result := DB.Where("username = ?", "test").First(&user)
	if result.RowsAffected == 0 {
		// Hash password sebelum disimpan
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		tanggal := "2000-01-01"

		// Parsing string menjadi time.Time
		parsedDate, err := time.Parse("2006-01-02", tanggal) // Gunakan format layout "2006-01-02"
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}

		// Buat user baru
		newUser := &entity.User{
			Nama:         "Admin",
			Alamat:       "Admin",
			Username:     "test",
			Email:        "M7K9O@example.com",
			Password:     string(hashedPassword),
			Role:         "ADMIN",
			TanggalLahir: parsedDate,
			JenisKelamin: "Laki-laki",
			NomorTelepon: "1234567890",
		}

		if err := DB.Create(&newUser).Error; err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
		log.Println("Admin user created successfully")
	} else {
		log.Println("Admin user already exists")
	}
}
