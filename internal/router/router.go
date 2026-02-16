package router

import (
	"github.com/gin-gonic/gin"

	"github.com/abhay786-20/fraud-auth-service/internal/config"
	"github.com/abhay786-20/fraud-auth-service/internal/handler"
	"github.com/abhay786-20/fraud-auth-service/internal/middleware"
	"github.com/abhay786-20/fraud-auth-service/pkg/logger"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(
	log *logger.Logger,
	cfg *config.Config,
	authHandler *handler.AuthHandler,
	healthHandler *handler.HealthHandler,
) *Router {

	gin.SetMode(cfg.Server.GinMode)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger(log))

	// Health check
	engine.GET("/health", healthHandler.Check)

	// Auth routes
	auth := engine.Group("/api/v1/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
	}

	log.Info("Router initialized")

	return &Router{
		Engine: engine,
	}
}
