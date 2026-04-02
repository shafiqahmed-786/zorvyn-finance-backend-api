package services

import (
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