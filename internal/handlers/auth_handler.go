package handlers

import (
    "finance-backend/internal/config"
    "finance-backend/internal/dto"
    "finance-backend/internal/models"
    "finance-backend/internal/services"
    "finance-backend/internal/utils"

    "github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
    Service *services.AuthService
    Config  *config.Config
}

// Register godoc
// @Summary Register a new user
// @Description Register user with role
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body dto.RegisterRequest true "Register payload"
// @Success 200 {object} models.User
// @Router /api/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
    var req dto.RegisterRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
    }

    hashed, _ := utils.HashPassword(req.Password)

    user := models.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: hashed,
        Role:     models.Role(req.Role),
        IsActive: true,
    }

    if user.Role == "" {
        user.Role = models.RoleViewer
    }

    err := h.Service.CreateUser(&user)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(user)
}

// Login godoc
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body dto.LoginRequest true "Login payload"
// @Success 200 {object} map[string]string
// @Router /api/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var req dto.LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
    }

    user, err := h.Service.GetByEmail(req.Email)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
    }

    if !utils.CheckPassword(req.Password, user.Password) {
        return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
    }

    token, _ := utils.GenerateJWT(user.ID.String(), string(user.Role), h.Config.JWTSecret)

    return c.JSON(fiber.Map{"token": token})
}