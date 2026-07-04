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

func setupTestRunRouter(h *endpoints.TestRunHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/test-runs/:id/cases", h.GetCasesByRun)
	r.POST("/test-runs/:id/cases", h.AddCaseToRun)
	r.DELETE("/test-runs/:id/cases/:test_case_id", h.RemoveCaseFromRun)
	return r
}

func TestGetCasesByRun_Success(t *testing.T) {
	mockRepo := new(mocks.TestRunRepository)
	mockRepo.On("GetCasesByRun", "run-1").Return([]entities.TestCase{
		{ID: "tc-1", Title: "Case One"},
		{ID: "tc-2", Title: "Case Two"},
	}, nil)

	r := setupTestRunRouter(endpoints.NewTestRunHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-runs/run-1/cases", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body []entities.TestCase
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Len(t, body, 2)

	mockRepo.AssertExpectations(t)
}

func TestGetCasesByRun_Error(t *testing.T) {
	mockRepo := new(mocks.TestRunRepository)
	mockRepo.On("GetCasesByRun", "bad-run").Return(nil, assert.AnError)

	r := setupTestRunRouter(endpoints.NewTestRunHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-runs/bad-run/cases", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestAddCaseToRun_Success(t *testing.T) {
	mockRepo := new(mocks.TestRunRepository)
	mockRepo.On("AddCaseToRun", "run-1", "tc-1").Return(nil)

	r := setupTestRunRouter(endpoints.NewTestRunHandler(mockRepo))

	body := map[string]string{"test_case_id": "tc-1"}
	payload, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/test-runs/run-1/cases", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestAddCaseToRun_Error(t *testing.T) {
	mockRepo := new(mocks.TestRunRepository)
	mockRepo.On("AddCaseToRun", "run-1", "tc-1").Return(assert.AnError)

	r := setupTestRunRouter(endpoints.NewTestRunHandler(mockRepo))

	body := map[string]string{"test_case_id": "tc-1"}
	payload, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/test-runs/run-1/cases", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestAddCaseToRun_BadRequest(t *testing.T) {
	mockRepo := new(mocks.TestRunRepository)

	r := setupTestRunRouter(endpoints.NewTestRunHandler(mockRepo))

	req := httptest.NewRequest(http.MethodPost, "/test-runs/run-1/cases", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRemoveCaseFromRun_Success(t *testing.T) {
	mockRepo := new(mocks.TestRunRepository)
	mockRepo.On("RemoveCaseFromRun", "run-1", "tc-1").Return(nil)

	r := setupTestRunRouter(endpoints.NewTestRunHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/test-runs/run-1/cases/tc-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestRemoveCaseFromRun_Error(t *testing.T) {
	mockRepo := new(mocks.TestRunRepository)
	mockRepo.On("RemoveCaseFromRun", "run-1", "tc-1").Return(assert.AnError)

	r := setupTestRunRouter(endpoints.NewTestRunHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/test-runs/run-1/cases/tc-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}
