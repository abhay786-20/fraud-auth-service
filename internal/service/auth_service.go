package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/abhay786-20/fraud-auth-service/internal/models"
	"github.com/abhay786-20/fraud-auth-service/internal/repository"
)


type AuthService struct {
	userRepo repository.UserRepository
	jwtSecret string
}

func NewAuthService(
	userRepo repository.UserRepository,
	jwtSecret string,
) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtSecret: jwtSecret,
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

