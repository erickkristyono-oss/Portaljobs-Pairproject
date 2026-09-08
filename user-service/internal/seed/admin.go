package seed

import (
	"log"

	"user-service/internal/domain/constant"
	"user-service/internal/repository/postgres/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	var count int64

	err := db.
		Model(&model.UserModel{}).
		Where("email = ?", "admin@gmail.com").
		Count(&count).Error

	if err != nil {
		log.Fatal("failed to check admin:", err)
	}

	// Admin sudah ada
	if count > 0 {
		log.Println("admin already exists")
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Fatal("failed to hash admin password:", err)
	}

	// Data admin
	admin := model.UserModel{
		Nama:     "Administrator",
		Email:    "admin@gmail.com",
		Password: string(hashedPassword),
		Role:     string(constant.RoleAdmin),
		Status:   "active",
	}

	// Insert admin
	if err := db.Create(&admin).Error; err != nil {
		log.Fatal("failed to seed admin:", err)
	}

	log.Println("admin seeded successfully")
}
