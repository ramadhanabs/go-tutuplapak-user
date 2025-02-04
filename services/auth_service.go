package services

import (
	"database/sql"
	"errors"
	"fmt"
	"go-tutuplapak-user/config"
	"go-tutuplapak-user/models"
	"go-tutuplapak-user/repositories"
	"go-tutuplapak-user/utils"
)

type AuthService interface {
	LoginWithEmail(email, password string) (*models.User, string, error)
	LoginWithPhone(phone, password string) (*models.User, string, error)
	RegisterWithEmail(email, password string) (*models.User, string, error)
	RegisterWithPhone(phone, password string) (*models.User, string, error)
}

type authService struct {
	userRepo repositories.UserRepository
	cfg      config.Config
}

func NewAuthService(userRepo repositories.UserRepository, cfg config.Config) AuthService {
	return &authService{userRepo: userRepo, cfg: cfg}
}

func (s *authService) LoginWithEmail(email, password string) (*models.User, string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", fmt.Errorf("%w: %v", utils.ErrInternal, err)
	}

	if user == nil {
		return nil, "", errors.New("email not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(email, s.cfg.JWTSecret, s.cfg.JWTExpiryHours)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *authService) LoginWithPhone(phone, password string) (*models.User, string, error) {
	user, err := s.userRepo.FindByPhone(phone)
	if err != nil {
		return nil, "", fmt.Errorf("%w: %v", utils.ErrInternal, err)
	}

	if user == nil {
		return nil, "", errors.New("phone not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(phone, s.cfg.JWTSecret, s.cfg.JWTExpiryHours)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *authService) RegisterWithEmail(email, password string) (*models.User, string, error) {
	exists, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, "", fmt.Errorf("%w: %v", utils.ErrInternal, err)
	}
	if exists != nil {
		return nil, "", errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, "", err
	}

	user := &models.User{
		Email:    sql.NullString{String: email, Valid: email != ""},
		Password: hashedPassword,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, "", err
	}

	token, err := utils.GenerateJWT(email, s.cfg.JWTSecret, s.cfg.JWTExpiryHours)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *authService) RegisterWithPhone(phone, password string) (*models.User, string, error) {

	if !utils.IsValidPhoneNumber(phone) {
		return nil, "", errors.New("phone number must start with '+' and be followed by digits")
	}

	exists, err := s.userRepo.FindByPhone(phone)
	if err != nil {
		return nil, "", fmt.Errorf("%w: %v", utils.ErrInternal, err)
	}
	if exists != nil {
		return nil, "", errors.New("phone already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, "", err
	}

	user := &models.User{
		Phone:    sql.NullString{String: phone, Valid: phone != ""},
		Password: hashedPassword,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, "", err
	}

	token, err := utils.GenerateJWT(phone, s.cfg.JWTSecret, s.cfg.JWTExpiryHours)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
