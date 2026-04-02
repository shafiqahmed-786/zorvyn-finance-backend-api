package services

import (
	"time"

	"finance-backend/internal/dto"
	"finance-backend/internal/models"

	"gorm.io/gorm"
)

type RecordService struct {
    DB *gorm.DB
}

func (s *RecordService) Create(record *models.FinancialRecord) error {
    return s.DB.Create(record).Error
}

func (s *RecordService) GetAll(records *[]models.FinancialRecord, page, limit int, filters map[string]string) error {
    query := s.DB.Model(&models.FinancialRecord{})

    if v := filters["type"]; v != "" {
        query = query.Where("type = ?", v)
    }
    if v := filters["category"]; v != "" {
        query = query.Where("category = ?", v)
    }

    offset := (page - 1) * limit
    return query.Offset(offset).Limit(limit).Order("date desc").Find(records).Error
}

func (s *RecordService) SoftDelete(id string) error {
    return s.DB.Delete(&models.FinancialRecord{}, "id = ?", id).Error
}

func (s *RecordService) CreateAudit(log *models.AuditLog) error {
    return s.DB.Create(log).Error
}

func (s *RecordService) Update(id string, req dto.UpdateRecordRequest) (*models.FinancialRecord, error) {
	var record models.FinancialRecord

	if err := s.DB.First(&record, "id = ?", id).Error; err != nil {
		return nil, err
	}

	record.Amount = req.Amount
	record.Type = models.RecordType(req.Type)
	record.Category = req.Category
	record.Notes = req.Notes

	if req.Date != "" {
		parsedDate, err := time.Parse("2006-01-02", req.Date)
		if err == nil {
			record.Date = parsedDate
		}
	}

	if err := s.DB.Save(&record).Error; err != nil {
		return nil, err
	}

	return &record, nil
}