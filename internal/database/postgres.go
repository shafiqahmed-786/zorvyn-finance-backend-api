package database

import (
    "fmt"
    "log"

    "finance-backend/internal/config"
    "finance-backend/internal/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.DBHost,
        cfg.DBUser,
        cfg.DBPass,
        cfg.DBName,
        cfg.DBPort,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database: ", err)
    }

    err = db.AutoMigrate(
	&models.User{},
	&models.FinancialRecord{},
	&models.AuditLog{},
    )
    if err != nil {
	log.Fatal("migration failed: ", err)
    }

    return db
}