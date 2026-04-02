package services

import (
	"finance-backend/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := s.DB.Order("created_at desc").Find(&users).Error
	return users, err
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := s.DB.First(&user, "id = ?", id).Error
	return &user, err
}

func (s *UserService) UpdateStatus(id string, isActive bool) error {
	return s.DB.Model(&models.User{}).
		Where("id = ?", id).
		Update("is_active", isActive).Error
}