package main

import (
	"log"

	"portaljob/internal/config"
	"portaljob/internal/database"
	"portaljob/internal/domain"
	"portaljob/internal/handler"
	"portaljob/internal/repository/postgres"
	"portaljob/internal/router"
	"portaljob/internal/usecase"
	"portaljob/pkg/jwt"
	"portaljob/pkg/notification"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatalf("database: %v", err)
	}

	jwtManager := jwt.NewManager(cfg.JWTSecret, cfg.JWTExpireHours)

	// notifier: real WA (Fonnte) if configured, otherwise a logging stub
	var notifier domain.Notifier
	if cfg.WAProvider == "fonnte" && cfg.FonnteToken != "" {
		notifier = notification.NewFonnte(cfg.FonnteToken, cfg.FonnteAPIURL)
		log.Println("notifier: fonnte")
	} else {
		notifier = notification.NewStub()
		log.Println("notifier: stub (logs only)")
	}

	// repositories
	userRepo := postgres.NewUserRepository(db)
	companyRepo := postgres.NewCompanyRepository(db)
	profileRepo := postgres.NewProfileRepository(db)
	skillRepo := postgres.NewSkillRepository(db)
	portfolioRepo := postgres.NewPortfolioRepository(db)
	jobRepo := postgres.NewJobRepository(db)
	appRepo := postgres.NewApplicationRepository(db)
	reportRepo := postgres.NewReportRepository(db)

	// usecases
	authUC := usecase.NewAuthUsecase(userRepo, jwtManager)
	companyUC := usecase.NewCompanyUsecase(companyRepo)
	profileUC := usecase.NewProfileUsecase(profileRepo)
	skillUC := usecase.NewSkillUsecase(skillRepo)
	portfolioUC := usecase.NewPortfolioUsecase(portfolioRepo)
	jobUC := usecase.NewJobUsecase(jobRepo, companyRepo)
	appUC := usecase.NewApplicationUsecase(appRepo, jobRepo, companyRepo, profileRepo, notifier)
	reportUC := usecase.NewReportUsecase(reportRepo)
	adminUC := usecase.NewAdminUsecase(userRepo, reportRepo)

	// handlers
	handlers := router.Handlers{
		Auth:        handler.NewAuthHandler(authUC),
		Job:         handler.NewJobHandler(jobUC),
		Company:     handler.NewCompanyHandler(companyUC),
		Profile:     handler.NewProfileHandler(profileUC),
		Skill:       handler.NewSkillHandler(skillUC),
		Portfolio:   handler.NewPortfolioHandler(portfolioUC),
		Application: handler.NewApplicationHandler(appUC),
		Report:      handler.NewReportHandler(reportUC),
		Admin:       handler.NewAdminHandler(adminUC),
	}

	r := router.New(handlers, jwtManager, cfg)

	log.Printf("server running on :%s (swagger: http://localhost:%s/swagger)", cfg.AppPort, cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("server: %v", err)
	}
}
