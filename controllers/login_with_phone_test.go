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

func TestLoginWithPhone(t *testing.T) {
	mockAuthServiceMock := new(services.AuthServiceMock)
	controller := controllers.NewAuthController(mockAuthServiceMock)

	router := utils.SetupRouter()
	router.POST("/v1/login/phone", controller.LoginWithPhone)

	t.Run("200 OK - Existing User", func(t *testing.T) {
		reqBody := map[string]string{"phone": "+6289898874", "password": "asdfasdf"}
		body, _ := json.Marshal(reqBody)

		mockAuthServiceMock.On("LoginWithPhone", "+6289898874", "asdfasdf").
			Return(&models.User{Phone: utils.NewNullableString("+6289898874")}, "token123", nil)

		req := httptest.NewRequest(http.MethodPost, "/v1/login/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		expectedResponse := `{"phone":"+6289898874", "email":"", "token":"token123"}`
		assert.JSONEq(t, expectedResponse, resp.Body.String())
	})

	t.Run("400 Bad Request - Validation Error Should Be String", func(t *testing.T) {
		reqBody := struct {
			Phone    int    `json:"phone"`
			Password string `json:"password"`
		}{
			Phone:    1,
			Password: "asdf",
		}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/v1/login/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("400 Bad Request - Validation Error: Required", func(t *testing.T) {
		reqBody := map[string]string{"phone": "", "password": "asdf"}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/v1/login/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("400 Bad Request - Validation Error: Should Be With + Prefix", func(t *testing.T) {
		reqBody := map[string]string{"phone": "4563099074", "password": "asdf"}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/v1/login/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("404 Not Found - Phone Not Found", func(t *testing.T) {
		reqBody := map[string]string{"phone": "+6743656478", "password": "asdfasdf"}
		body, _ := json.Marshal(reqBody)

		mockAuthServiceMock.On("LoginWithPhone", "+6743656478", "asdfasdf").
			Return(nil, "", errors.New("phone not found"))

		req := httptest.NewRequest(http.MethodPost, "/v1/login/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
		expectedResponse := `{"error":"phone not found"}`
		assert.JSONEq(t, expectedResponse, resp.Body.String())
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		mockAuthServiceMock := new(services.AuthServiceMock)
		controller := controllers.NewAuthController(mockAuthServiceMock)

		router := utils.SetupRouter()
		router.POST("/v1/login/phone", controller.LoginWithPhone)

		reqBody := map[string]string{"phone": "+6743656478", "password": "asdfasdf"}
		body, _ := json.Marshal(reqBody)

		mockAuthServiceMock.On("LoginWithPhone", "+6743656478", "asdfasdf").
			Return(nil, "", utils.ErrInternal)

		req := httptest.NewRequest(http.MethodPost, "/v1/login/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)

		expectedResponse := `{"error":"internal server error"}`
		assert.JSONEq(t, expectedResponse, resp.Body.String())
	})
}
