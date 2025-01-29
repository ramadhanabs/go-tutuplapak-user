package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-tutuplapak-user/controllers"
	"go-tutuplapak-user/models"
	"go-tutuplapak-user/services"
	"go-tutuplapak-user/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterWithEmail(t *testing.T) {
	mockAuthServiceMock := new(services.AuthServiceMock)
	controller := controllers.NewAuthController(mockAuthServiceMock)

	router := utils.SetupRouter()
	router.POST("/v1/register/email", controller.RegisterWithEmail)

	t.Run("201 OK - No Existing User", func(t *testing.T) {
		reqBody := map[string]string{"email": "name@name.com", "password": "asdfasdf"}
		body, _ := json.Marshal(reqBody)

		mockAuthServiceMock.On("RegisterWithEmail", "name@name.com", "asdfasdf").
			Return(&models.User{Email: utils.NewNullableString("name@name.com")}, "token123", nil)

		req := httptest.NewRequest(http.MethodPost, "/v1/register/email", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
		expectedResponse := `{"email":"name@name.com", "phone":"", "token":"token123"}`
		assert.JSONEq(t, expectedResponse, resp.Body.String())
	})

	t.Run("400 Bad Request - Validation Error Should Be String", func(t *testing.T) {
		reqBody := struct {
			Email    int    `json:"email"`
			Password string `json:"password"`
		}{
			Email:    1,
			Password: "asdf",
		}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/v1/register/email", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("400 Bad Request - Validation Error: Required", func(t *testing.T) {
		reqBody := map[string]string{"email": "", "password": "asdf"}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/v1/register/email", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("400 Bad Request - Validation Error: Should Be Email", func(t *testing.T) {
		reqBody := map[string]string{"email": "invalid-email", "password": "asdf"}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/v1/register/email", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("409 OK - User Already Exist", func(t *testing.T) {
		mockAuthServiceMock := new(services.AuthServiceMock)
		controller := controllers.NewAuthController(mockAuthServiceMock)

		router := utils.SetupRouter()
		router.POST("/v1/register/email", controller.RegisterWithEmail)

		reqBody := map[string]string{"email": "name@name.com", "password": "asdfasdf"}
		body, _ := json.Marshal(reqBody)

		mockAuthServiceMock.On("RegisterWithEmail", "name@name.com", "asdfasdf").
			Return(nil, "", errors.New("email already exists")).Once()

		req := httptest.NewRequest(http.MethodPost, "/v1/register/email", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusConflict, resp.Code)
		expectedResponse := `{"error":"email already exists"}`
		assert.JSONEq(t, expectedResponse, resp.Body.String())
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		mockAuthServiceMock := new(services.AuthServiceMock)
		controller := controllers.NewAuthController(mockAuthServiceMock)

		router := utils.SetupRouter()
		router.POST("/v1/register/email", controller.RegisterWithEmail)

		reqBody := map[string]string{"email": "name@name.com", "password": "asdfasdf"}
		body, _ := json.Marshal(reqBody)

		mockAuthServiceMock.On("RegisterWithEmail", "name@name.com", "asdfasdf").
			Return(nil, "", utils.ErrInternal)

		req := httptest.NewRequest(http.MethodPost, "/v1/register/email", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)

		expectedResponse := `{"error":"internal server error"}`
		assert.JSONEq(t, expectedResponse, resp.Body.String())
	})
}
