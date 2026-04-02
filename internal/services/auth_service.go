package services

import (
    "finance-backend/internal/models"

    "gorm.io/gorm"
)

type AuthService struct {
    DB *gorm.DB
}

func (s *AuthService) CreateUser(user *models.User) error {
    return s.DB.Create(user).Error
}

func (s *AuthService) GetByEmail(email string) (*models.User, error) {
    var user models.User
    err := s.DB.Where("email = ?", email).First(&user).Error
    return &user, err
}