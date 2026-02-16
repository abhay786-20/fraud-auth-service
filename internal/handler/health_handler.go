package handler

import (
	"net/http"

	"github.com/abhay786-20/fraud-auth-service/internal/db"
	"github.com/abhay786-20/fraud-auth-service/internal/dto"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	DB *db.Postgres
}

func NewHealthHandler(pg *db.Postgres) *HealthHandler {
	return &HealthHandler{
		DB: pg,
	}
}

func (h *HealthHandler) Check(c *gin.Context) {
	services := make(map[string]string)
	overallStatus := "healthy"

	// Check database
	if err := h.DB.Ping(); err != nil {
		services["database"] = "unhealthy"
		overallStatus = "unhealthy"
	} else {
		services["database"] = "healthy"
	}

	statusCode := http.StatusOK
	if overallStatus == "unhealthy" {
		statusCode = http.StatusServiceUnavailable
	}

	c.JSON(statusCode, dto.HealthResponse{
		Status:   overallStatus,
		Services: services,
	})
}
