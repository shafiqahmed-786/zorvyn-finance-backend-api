package services

import (
	"finance-backend/internal/dto"
	"finance-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

type DashboardService struct {
	DB *gorm.DB
}

func (s *DashboardService) GetSummary() (*dto.DashboardSummaryResponse, error) {
	var totalIncome float64
	var totalExpense float64

	s.DB.Model(&models.FinancialRecord{}).
		Where("type = ?", "income").
		Select("COALESCE(SUM(amount),0)").
		Scan(&totalIncome)

	s.DB.Model(&models.FinancialRecord{}).
		Where("type = ?", "expense").
		Select("COALESCE(SUM(amount),0)").
		Scan(&totalExpense)

	var recentCount int64
	last7Days := time.Now().AddDate(0, 0, -7)
	s.DB.Model(&models.FinancialRecord{}).
		Where("created_at >= ?", last7Days).
		Count(&recentCount)

	return &dto.DashboardSummaryResponse{
		TotalIncome:   totalIncome,
		TotalExpense:  totalExpense,
		NetBalance:    totalIncome - totalExpense,
		CategoryTotals: map[string]float64{},
		RecentCount:   recentCount,
	}, nil
}

func (s *DashboardService) GetMonthlyTrend() ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	return result, nil
}

func (s *DashboardService) GetRecentActivity(limit int) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := s.DB.Order("created_at desc").Limit(limit).Find(&logs).Error
	return logs, err
}