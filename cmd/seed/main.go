package main

import (
    "finance-backend/internal/config"
    "finance-backend/internal/database"
    "finance-backend/internal/models"
)

func main() {
    cfg := config.Load()
    db := database.Connect(cfg)

    db.Create(&models.FinancialRecord{
        Amount: 50000,
        Type: "income",
        Category: "salary",
    })

    db.Create(&models.FinancialRecord{
        Amount: 3000,
        Type: "expense",
        Category: "food",
    })
}