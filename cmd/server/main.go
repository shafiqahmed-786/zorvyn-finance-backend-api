// @title Finance Backend API
// @version 1.0
// @description Zorvyn Backend Developer Internship Assignment - Production-grade fintech backend with RBAC, analytics, audit logs, and Swagger docs
// @host localhost:8082
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"log"

	"finance-backend/internal/config"
	"finance-backend/internal/database"
	"finance-backend/internal/routes"

	_ "finance-backend/docs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	app := fiber.New()
	routes.Setup(app, db, cfg)

	log.Fatal(app.Listen(":" + cfg.Port))
}