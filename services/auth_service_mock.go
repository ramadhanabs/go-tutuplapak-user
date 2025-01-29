package services

import (
	"go-tutuplapak-user/models"

	"github.com/stretchr/testify/mock"
)

type AuthServiceMock struct {
	mock.Mock
}

func (m *AuthServiceMock) LoginWithEmail(email, password string) (*models.User, string, error) {
	args := m.Called(email, password)
	user, _ := args.Get(0).(*models.User)
	token, _ := args.Get(1).(string)
	return user, token, args.Error(2)
}

func (m *AuthServiceMock) LoginWithPhone(phone, password string) (*models.User, string, error) {
	args := m.Called(phone, password)
	user, _ := args.Get(0).(*models.User)
	token, _ := args.Get(1).(string)
	return user, token, args.Error(2)
}

func (m *AuthServiceMock) RegisterWithEmail(email, password string) (*models.User, string, error) {
	args := m.Called(email, password)
	user, _ := args.Get(0).(*models.User)
	token, _ := args.Get(1).(string)
	return user, token, args.Error(2)
}

func (m *AuthServiceMock) RegisterWithPhone(phone, password string) (*models.User, string, error) {
	args := m.Called(phone, password)
	user, _ := args.Get(0).(*models.User)
	token, _ := args.Get(1).(string)
	return user, token, args.Error(2)
}
