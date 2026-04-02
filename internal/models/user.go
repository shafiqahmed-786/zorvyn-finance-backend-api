package models

import (
    "time"

    "github.com/google/uuid"
)

type Role string

const (
    RoleViewer Role = "viewer"
    RoleAnalyst Role = "analyst"
    RoleAdmin Role = "admin"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Name      string    `gorm:"not null" json:"name"`
    Email     string    `gorm:"uniqueIndex;not null" json:"email"`
    Password  string    `gorm:"not null" json:"-"`
    Role      Role      `gorm:"type:varchar(20);default:'viewer'" json:"role"`
    IsActive  bool      `gorm:"default:true" json:"is_active"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}