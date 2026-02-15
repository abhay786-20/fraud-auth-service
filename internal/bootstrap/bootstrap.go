package bootstrap

import (
	"github.com/abhay786-20/fraud-auth-service/internal/config"
	"github.com/abhay786-20/fraud-auth-service/internal/db"
	"github.com/abhay786-20/fraud-auth-service/internal/router"
	"github.com/abhay786-20/fraud-auth-service/internal/service"
	"github.com/abhay786-20/fraud-auth-service/internal/repository"
	"github.com/abhay786-20/fraud-auth-service/pkg/env"
	"github.com/abhay786-20/fraud-auth-service/pkg/logger"
)

type Application struct {
	Config *config.Config
	Logger *logger.Logger
	DB     *db.Postgres
	Router *router.Router
}

func NewApplication() (*Application, error) {

	// 1️⃣ Load Environment from .env file and validate required vars
	environment, err := env.New()
	if err != nil {
		return nil, err
	}

	// 2️⃣ Load Config
	cfg := config.LoadConfig(environment)

	// 3️⃣ Logger
	log := logger.New()
	log.Info("Starting Fraud Auth Service")

	// 4️⃣ Database
	pg, err := db.NewPostgres(cfg.Database)
	if err != nil {
		log.Error("Database connection failed")
		return nil, err
	}
	log.Info("Database connected")

	// Repository - User
	userRepo := repository.NewPostgresUserRepository(pg.DB, log)

	// Service - Auth
	authService := service.NewAuthService(userRepo, cfg.Auth.JWTSecret)		

	// 5️⃣ Router
	r := router.NewRouter(log, cfg, pg, authService)

	return &Application{
		Config: cfg,
		Logger: log,
		DB:     pg,
		Router: r,
	}, nil
}
