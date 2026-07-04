package endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/xurceo/plain-tcms/endpoints"
	"github.com/xurceo/plain-tcms/entities"
	"github.com/xurceo/plain-tcms/tests/mocks"
)

func setupAuthRouter(h *endpoints.AuthHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)
	return r
}

func TestRegister_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	req := entities.RegisterRequest{Email: "test@test.com", Password: "pass123"}
	mockRepo.On("Register", "test@test.com", "pass123").Return(entities.User{
		ID: "u-1", Email: "test@test.com",
	}, nil)

	r := setupAuthRouter(endpoints.NewAuthHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entities.AuthResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "test@test.com", resp.User.Email)

	mockRepo.AssertExpectations(t)
}

func TestRegister_BadRequest(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	r := setupAuthRouter(endpoints.NewAuthHandler(mockRepo))

	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRegister_Error(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	req := entities.RegisterRequest{Email: "test@test.com", Password: "pass123"}
	mockRepo.On("Register", "test@test.com", "pass123").Return(entities.User{}, assert.AnError)

	r := setupAuthRouter(endpoints.NewAuthHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	req := entities.LoginRequest{Email: "test@test.com", Password: "pass123"}
	mockRepo.On("Login", "test@test.com", "pass123").Return(entities.User{
		ID: "u-1", Email: "test@test.com",
	}, nil)

	r := setupAuthRouter(endpoints.NewAuthHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp entities.AuthResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "test@test.com", resp.User.Email)

	mockRepo.AssertExpectations(t)
}

func TestLogin_BadRequest(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	r := setupAuthRouter(endpoints.NewAuthHandler(mockRepo))

	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	req := entities.LoginRequest{Email: "test@test.com", Password: "wrong"}
	mockRepo.On("Login", "test@test.com", "wrong").Return(entities.User{}, assert.AnError)

	r := setupAuthRouter(endpoints.NewAuthHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	mockRepo.AssertExpectations(t)
}
