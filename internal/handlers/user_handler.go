package handlers

import (
    "finance-backend/internal/models"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type UserHandler struct {
    DB *gorm.DB
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
    var users []models.User
    h.DB.Find(&users)
    return c.JSON(users)
}