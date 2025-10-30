package seed

import (
	"log"

	"github.com/google/uuid"
	"github.com/hfishere/wpu-project-management/config"
	"github.com/hfishere/wpu-project-management/models"
	"github.com/hfishere/wpu-project-management/utils"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name:     "Super Admin",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
		PublicID: uuid.New(),
	}

	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed to seed admin", err)
	} else {
		log.Println("Admin user seeded.")
	}
}
