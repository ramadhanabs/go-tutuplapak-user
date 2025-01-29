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

func TestRegisterWithPhone(t *testing.T) {
	mockAuthServiceMock := new(services.AuthServiceMock)
	controller := controllers.NewAuthController(mockAuthServiceMock)

	router := utils.SetupRouter()
	router.POST("/v1/register/phone", controller.RegisterWithPhone)

	t.Run("201 OK - No Existing User", func(t *testing.T) {
		reqBody := map[string]string{"phone": "+548877653745", "password": "asdfasdf"}
		body, _ := json.Marshal(reqBody)

		mockAuthServiceMock.On("RegisterWithPhone", "+548877653745", "asdfasdf").
			Return(&models.User{Phone: utils.NewNullableString("+548877653745")}, "token123", nil)

		req := httptest.NewRequest(http.MethodPost, "/v1/register/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
		expectedResponse := `{"phone":"+548877653745", "email":"", "token":"token123"}`
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

		req := httptest.NewRequest(http.MethodPost, "/v1/register/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("400 Bad Request - Validation Error: Required", func(t *testing.T) {
		reqBody := map[string]string{"phone": "", "password": "asdf"}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/v1/register/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("400 Bad Request - Validation Error: Should Be With + Prefix", func(t *testing.T) {
		reqBody := map[string]string{"phone": "67675899", "password": "asdf"}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/v1/register/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})

	t.Run("409 OK - User Already Exist", func(t *testing.T) {
		mockAuthServiceMock := new(services.AuthServiceMock)
		controller := controllers.NewAuthController(mockAuthServiceMock)

		router := utils.SetupRouter()
		router.POST("/v1/register/phone", controller.RegisterWithPhone)

		reqBody := map[string]string{"phone": "+67675899", "password": "asdfasdf"}
		body, _ := json.Marshal(reqBody)

		mockAuthServiceMock.On("RegisterWithPhone", "+67675899", "asdfasdf").
			Return(nil, "", errors.New("phone already exists")).Once()

		req := httptest.NewRequest(http.MethodPost, "/v1/register/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusConflict, resp.Code)
		expectedResponse := `{"error":"phone already exists"}`
		assert.JSONEq(t, expectedResponse, resp.Body.String())
	})

	t.Run("500 Internal Server Error", func(t *testing.T) {
		mockAuthServiceMock := new(services.AuthServiceMock)
		controller := controllers.NewAuthController(mockAuthServiceMock)

		router := utils.SetupRouter()
		router.POST("/v1/register/phone", controller.RegisterWithPhone)

		reqBody := map[string]string{"phone": "+67675899", "password": "asdfasdf"}
		body, _ := json.Marshal(reqBody)

		mockAuthServiceMock.On("RegisterWithPhone", "+67675899", "asdfasdf").
			Return(nil, "", utils.ErrInternal)

		req := httptest.NewRequest(http.MethodPost, "/v1/register/phone", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)

		expectedResponse := `{"error":"internal server error"}`
		assert.JSONEq(t, expectedResponse, resp.Body.String())
	})
}
