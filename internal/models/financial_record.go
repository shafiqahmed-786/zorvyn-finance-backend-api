package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type RecordType string

const (
    TypeIncome  RecordType = "income"
    TypeExpense RecordType = "expense"
)

type FinancialRecord struct {
    ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    UserID    uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
    Amount    float64        `gorm:"not null" json:"amount"`
    Type      RecordType     `gorm:"type:varchar(20);not null" json:"type"`
    Category  string         `gorm:"not null" json:"category"`
    Date      time.Time      `gorm:"not null" json:"date"`
    Notes     string         `json:"notes"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}