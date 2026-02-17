package service

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/abhay786-20/fraud-auth-service/internal/models"
	"github.com/abhay786-20/fraud-auth-service/internal/repository"
	"github.com/abhay786-20/fraud-auth-service/pkg/utils"
)


type AuthService struct {
	userRepo  repository.UserRepository
	jwtSecret string
	tokenTTL  time.Duration
}

func NewAuthService(
	userRepo repository.UserRepository,
	jwtSecret string,
	tokenTTL time.Duration,
) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		tokenTTL:  tokenTTL,
	}
}

func (s *AuthService) Signup(email, password string) (*models.User, error) {

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	// Save to DB
	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}


func (s *AuthService) Login(email, password string) (*models.User, error) {

	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) GenerateToken(user *models.User) (string, error) {
	return utils.GenerateToken(
		user.ID,
		user.Email,
		s.jwtSecret,
		s.tokenTTL,
	)
}

