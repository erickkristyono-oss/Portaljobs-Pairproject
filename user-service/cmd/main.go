package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"

	"user-service/config"
	"user-service/internal/handler"
	"user-service/internal/repository/postgres"
	"user-service/internal/router"
	"user-service/internal/seed"

	usecase_portfolio "user-service/internal/usecase/portfolio"
	usecase_profile "user-service/internal/usecase/profile"
	usecase_skill "user-service/internal/usecase/skill"
	usecase_user "user-service/internal/usecase/user"
)

func main() {

	// Load environment variable
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found")
	}

	// Database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	log.Println("database connected successfully")

	// Seed admin
	seed.SeedAdmin(db)

	// Repository
	userRepository := postgres.NewUserRepository(db)
	profileRepository := postgres.NewProfileRepository(db)
	skillRepository := postgres.NewSkillRepository(db)
	portfolioRepository := postgres.NewPortfolioRepository(db)

	// Usecase
	userUsecase := usecase_user.NewUserUsecase(userRepository)
	profileUsecase := usecase_profile.NewProfileUsecase(profileRepository)
	skillUsecase := usecase_skill.NewSkillUsecase(skillRepository)
	portfolioUsecase := usecase_portfolio.NewPortfolioUsecase(portfolioRepository)

	// Handler
	userHandler := handler.NewUserHandler(userUsecase)
	profileHandler := handler.NewProfileHandler(profileUsecase)
	skillHandler := handler.NewSkillHandler(skillUsecase)
	portfolioHandler := handler.NewPortfolioHandler(portfolioUsecase)

	// Echo
	e := echo.New()

	// Router
	router.RegisterRoutes(
		e,
		userHandler,
		profileHandler,
		skillHandler,
		portfolioHandler,
	)

	// Server
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8081"
	}

	log.Println("user-service running on http://localhost:" + port)

	if err := e.Start(":" + port); err != nil {
		log.Fatal("failed to start user-service:", err)
	}
}
