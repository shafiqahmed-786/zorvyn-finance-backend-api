package handlers

import (
	"finance-backend/internal/dto"
	"finance-backend/internal/models"
	"finance-backend/internal/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RecordHandler struct {
	Service *services.RecordService
}

// CreateRecord godoc
// @Summary Create financial record
// @Tags Records
// @Accept json
// @Produce json
// @Param payload body dto.CreateRecordRequest true "Record payload"
// @Success 200 {object} models.FinancialRecord
// @Router /api/records [post]
func (h *RecordHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateRecordRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid date format use YYYY-MM-DD"})
	}

	userID := c.Locals("user_id").(string)
	uid, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	record := models.FinancialRecord{
		UserID:   uid,
		Amount:   req.Amount,
		Type:     models.RecordType(req.Type),
		Category: req.Category,
		Date:     parsedDate,
		Notes:    req.Notes,
	}

	if err := h.Service.Create(&record); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	_ = h.Service.CreateAudit(&models.AuditLog{
		ActorID:      userID,
		Action:       "CREATE",
		ResourceType: "financial_record",
		ResourceID:   record.ID.String(),
	})

	return c.JSON(record)
}

// UpdateRecord godoc
// @Summary Update financial record
// @Tags Records
// @Accept json
// @Produce json
// @Param id path string true "Record ID"
// @Param payload body dto.UpdateRecordRequest true "Update payload"
// @Success 200 {object} models.FinancialRecord
// @Router /api/records/{id} [put]
func (h *RecordHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateRecordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	record, err := h.Service.Update(id, req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(record)
}

// GetAllRecords godoc
// @Summary Get all records
// @Tags Records
// @Produce json
// @Success 200 {array} models.FinancialRecord
// @Router /api/records [get]
func (h *RecordHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	filters := map[string]string{
		"type":     c.Query("type"),
		"category": c.Query("category"),
	}

	var records []models.FinancialRecord
	err := h.Service.GetAll(&records, page, limit, filters)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(records)
}

func (h *RecordHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.Service.SoftDelete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	_ = h.Service.CreateAudit(&models.AuditLog{
		ActorID:      c.Locals("user_id").(string),
		Action:       "DELETE",
		ResourceType: "financial_record",
		ResourceID:   id,
	})

	return c.JSON(fiber.Map{"message": "record soft deleted"})
}