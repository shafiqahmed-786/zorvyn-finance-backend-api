package routes

import (
	"finance-backend/internal/config"
	"finance-backend/internal/handlers"
	"finance-backend/internal/middleware"
	"finance-backend/internal/services"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	authService := &services.AuthService{DB: db}
	authHandler := &handlers.AuthHandler{
		Service: authService,
		Config:  cfg,
	}

	userHandler := &handlers.UserHandler{DB: db}

	recordService := &services.RecordService{DB: db}
	recordHandler := &handlers.RecordHandler{
		Service: recordService,
	}

	dashboardService := &services.DashboardService{DB: db}
	dashboardHandler := &handlers.DashboardHandler{
		Service: dashboardService,
	}

	// swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api")

	// auth routes
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// user routes
	users := api.Group("/users", middleware.Protected(cfg.JWTSecret))
	users.Get("/", middleware.RequireRoles("admin"), userHandler.GetUsers)

	// financial record routes
	records := api.Group("/records", middleware.Protected(cfg.JWTSecret))
	records.Get("/", middleware.RequireRoles("admin", "analyst", "viewer"), recordHandler.GetAll)
	records.Post("/", middleware.RequireRoles("admin"), recordHandler.Create)
	records.Delete("/:id", middleware.RequireRoles("admin"), recordHandler.Delete)

	// dashboard analytics routes
	dashboard := api.Group("/dashboard", middleware.Protected(cfg.JWTSecret))
	dashboard.Get("/summary", middleware.RequireRoles("admin", "analyst", "viewer"), dashboardHandler.Summary)
	dashboard.Get("/trends", middleware.RequireRoles("admin", "analyst"), dashboardHandler.Trends)
	dashboard.Get("/recent-activity", middleware.RequireRoles("admin", "analyst"), dashboardHandler.RecentActivity)
}