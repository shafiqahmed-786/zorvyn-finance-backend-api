package models

import (
    "time"

    "github.com/google/uuid"
)

type AuditLog struct {
    ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    ActorID      string    `json:"actor_id"`
    Action       string    `json:"action"`
    ResourceType string    `json:"resource_type"`
    ResourceID   string    `json:"resource_id"`
    CreatedAt    time.Time `json:"created_at"`
}