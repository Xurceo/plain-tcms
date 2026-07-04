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

func setupTestSuiteRouter(h *endpoints.TestSuiteHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/projects/:id/test-suites", h.GetTestSuitesByProject)
	r.POST("/projects/:id/test-suites", h.CreateTestSuite)
	r.GET("/test-suites/:id", h.GetTestSuiteByID)
	r.PUT("/test-suites/:id", h.UpdateTestSuite)
	r.DELETE("/test-suites/:id", h.DeleteTestSuite)
	return r
}

func TestGetTestSuitesByProject_Success(t *testing.T) {
	mockRepo := new(mocks.TestSuiteRepository)
	mockRepo.On("GetTestSuitesByProject", "proj-1").Return([]entities.TestSuite{
		{ID: "ts-1", Name: "Suite One"},
		{ID: "ts-2", Name: "Suite Two"},
	}, nil)

	r := setupTestSuiteRouter(endpoints.NewTestSuiteHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/projects/proj-1/test-suites", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body []entities.TestSuite
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Len(t, body, 2)

	mockRepo.AssertExpectations(t)
}

func TestGetTestSuitesByProject_Error(t *testing.T) {
	mockRepo := new(mocks.TestSuiteRepository)
	mockRepo.On("GetTestSuitesByProject", "bad-id").Return(nil, assert.AnError)

	r := setupTestSuiteRouter(endpoints.NewTestSuiteHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/projects/bad-id/test-suites", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateTestSuite_Success(t *testing.T) {
	mockRepo := new(mocks.TestSuiteRepository)
	req := entities.CreateTestSuiteRequest{Name: "New Suite"}
	mockRepo.On("CreateTestSuite", "proj-1", req).Return(entities.TestSuite{
		ID: "ts-3", Name: "New Suite",
	}, nil)

	r := setupTestSuiteRouter(endpoints.NewTestSuiteHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/projects/proj-1/test-suites", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entities.TestSuite
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "New Suite", resp.Name)

	mockRepo.AssertExpectations(t)
}

func TestCreateTestSuite_BadRequest(t *testing.T) {
	mockRepo := new(mocks.TestSuiteRepository)

	r := setupTestSuiteRouter(endpoints.NewTestSuiteHandler(mockRepo))

	req := httptest.NewRequest(http.MethodPost, "/projects/proj-1/test-suites", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTestSuiteByID_Success(t *testing.T) {
	mockRepo := new(mocks.TestSuiteRepository)
	mockRepo.On("GetTestSuiteByID", "ts-1").Return(entities.TestSuite{
		ID: "ts-1", Name: "Suite One",
	}, nil)

	r := setupTestSuiteRouter(endpoints.NewTestSuiteHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-suites/ts-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body entities.TestSuite
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Equal(t, "Suite One", body.Name)

	mockRepo.AssertExpectations(t)
}

func TestGetTestSuiteByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.TestSuiteRepository)
	mockRepo.On("GetTestSuiteByID", "999").Return(entities.TestSuite{}, assert.AnError)

	r := setupTestSuiteRouter(endpoints.NewTestSuiteHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-suites/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTestSuite_Success(t *testing.T) {
	mockRepo := new(mocks.TestSuiteRepository)
	req := entities.CreateTestSuiteRequest{Name: "Updated Suite"}
	mockRepo.On("UpdateTestSuite", "ts-1", req).Return(entities.TestSuite{
		ID: "ts-1", Name: "Updated Suite",
	}, nil)

	r := setupTestSuiteRouter(endpoints.NewTestSuiteHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/test-suites/ts-1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp entities.TestSuite
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "Updated Suite", resp.Name)

	mockRepo.AssertExpectations(t)
}

func TestUpdateTestSuite_Error(t *testing.T) {
	mockRepo := new(mocks.TestSuiteRepository)
	req := entities.CreateTestSuiteRequest{Name: "Updated Suite"}
	mockRepo.On("UpdateTestSuite", "ts-1", req).Return(entities.TestSuite{}, assert.AnError)

	r := setupTestSuiteRouter(endpoints.NewTestSuiteHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/test-suites/ts-1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTestSuite_Success(t *testing.T) {
	mockRepo := new(mocks.TestSuiteRepository)
	mockRepo.On("DeleteTestSuite", "ts-1").Return(nil)

	r := setupTestSuiteRouter(endpoints.NewTestSuiteHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/test-suites/ts-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTestSuite_Error(t *testing.T) {
	mockRepo := new(mocks.TestSuiteRepository)
	mockRepo.On("DeleteTestSuite", "ts-1").Return(assert.AnError)

	r := setupTestSuiteRouter(endpoints.NewTestSuiteHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/test-suites/ts-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}
