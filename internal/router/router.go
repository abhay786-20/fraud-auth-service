package router

import (
	"github.com/gin-gonic/gin"

	"github.com/abhay786-20/fraud-auth-service/internal/config"
	"github.com/abhay786-20/fraud-auth-service/internal/db"
	"github.com/abhay786-20/fraud-auth-service/internal/middleware"
	"github.com/abhay786-20/fraud-auth-service/internal/service"
	"github.com/abhay786-20/fraud-auth-service/pkg/logger"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(
	log *logger.Logger,
	cfg *config.Config,
	pg *db.Postgres,
	authService *service.AuthService,
) *Router {

	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger(log))

	// Health check route
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	log.Info("Router initialized")

	return &Router{
		Engine: engine,
	}
}
