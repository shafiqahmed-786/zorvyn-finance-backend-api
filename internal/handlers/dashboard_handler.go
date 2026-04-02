package handlers

import (
	"finance-backend/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	Service *services.DashboardService
}

// DashboardSummary godoc
// @Summary Get dashboard summary
// @Tags Dashboard
// @Produce json
// @Success 200 {object} dto.DashboardSummaryResponse
// @Router /api/dashboard/summary [get]
func (h *DashboardHandler) Summary(c *fiber.Ctx) error {
	data, err := h.Service.GetSummary()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

func (h *DashboardHandler) Trends(c *fiber.Ctx) error {
	data, err := h.Service.GetMonthlyTrend()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

func (h *DashboardHandler) RecentActivity(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "5"))

	data, err := h.Service.GetRecentActivity(limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}