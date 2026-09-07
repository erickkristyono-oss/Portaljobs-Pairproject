package seed

import (
	"errors"
	"log"

	"user-service/internal/domain/constant"
	"user-service/internal/repository/postgres/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	var existingUser model.UserModel

	err := db.
		Where("email = ?", "admin@gmail.com").
		First(&existingUser).
		Error

	// Admin sudah ada
	if err == nil {
		log.Println("admin already exists")
		return
	}

	// Error selain record not found
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatal("failed to check admin:", err)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Fatal("failed to hash admin password:", err)
	}

	// Buat admin
	admin := model.UserModel{
		Nama:     "Administrator",
		Email:    "admin@gmail.com",
		Password: string(hashedPassword),
		Role:     string(constant.RoleAdmin),
		Status:   "active",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatal("failed to seed admin:", err)
	}

	log.Println("admin seeded successfully")
}
