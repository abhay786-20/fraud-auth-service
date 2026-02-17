package handler

import (
	"net/http"

	"github.com/abhay786-20/fraud-auth-service/internal/dto"
	"github.com/abhay786-20/fraud-auth-service/internal/service"
	"github.com/abhay786-20/fraud-auth-service/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *service.AuthService
	Logger  *logger.Logger
}

func NewAuthHandler(
	service *service.AuthService,
	log *logger.Logger,
) *AuthHandler {
	return &AuthHandler{
		Service: service,
		Logger:  log,
	}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var req dto.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user, err := h.Service.Signup(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.SignupResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	user, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Error: "invalid credentials",
		})
		return
	}

	token, err := h.Service.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Token: token,
	})
}
