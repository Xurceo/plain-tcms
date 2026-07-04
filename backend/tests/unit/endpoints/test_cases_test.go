/*
 * // TCMS - Test Case Management System
 * // Copyright (C) 2026 Pavlo Shnal
 * //
 * // This program is free software: you can redistribute it and/or modify
 * // it under the terms of the GNU Affero General Public License as published
 * // by the Free Software Foundation, either version 3 of the License, or
 * // (at your option) any later version.
 * //
 * // This program is distributed in the hope that it will be useful,
 * // but WITHOUT ANY WARRANTY; without even the implied warranty of
 * // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * // GNU Affero General Public License for more details.
 * //
 * // You should have received a copy of the GNU Affero General Public License
 * // along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

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

func setupTestCaseRouter(h *endpoints.TestCaseHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/projects/:id/test-cases", h.GetTestCases)
	r.POST("/projects/:id/test-cases", h.CreateTestCase)
	r.GET("/test-cases/:id", h.GetTestCaseByID)
	r.PUT("/test-cases/:id", h.UpdateTestCase)
	r.DELETE("/test-cases/:id", h.DeleteTestCase)
	r.GET("/test-cases/:id/history", h.GetTestCaseHistory)
	return r
}

func TestGetTestCases_Success(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	mockRepo.On("GetTestCases", "proj-1").Return([]entities.TestCase{
		{ID: "tc-1", Title: "Case One"},
		{ID: "tc-2", Title: "Case Two"},
	}, nil)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/projects/proj-1/test-cases", nil)
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

func TestGetTestCases_Error(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	mockRepo.On("GetTestCases", "bad-id").Return(nil, assert.AnError)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/projects/bad-id/test-cases", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateTestCase_Success(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	req := entities.CreateTestCaseRequest{Title: "New Case"}
	mockRepo.On("CreateTestCase", "proj-1", req).Return(entities.TestCase{
		ID: "tc-3", Title: "New Case",
	}, nil)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/projects/proj-1/test-cases", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entities.TestCase
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "New Case", resp.Title)

	mockRepo.AssertExpectations(t)
}

func TestCreateTestCase_BadRequest(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	req := httptest.NewRequest(http.MethodPost, "/projects/proj-1/test-cases", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateTestCase_Error(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	req := entities.CreateTestCaseRequest{Title: "New Case"}
	mockRepo.On("CreateTestCase", "proj-1", req).Return(entities.TestCase{}, assert.AnError)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/projects/proj-1/test-cases", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetTestCaseByID_Success(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	mockRepo.On("GetTestCaseByID", "tc-1").Return(entities.TestCase{
		ID: "tc-1", Title: "Case One",
	}, nil)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-cases/tc-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body entities.TestCase
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Equal(t, "Case One", body.Title)

	mockRepo.AssertExpectations(t)
}

func TestGetTestCaseByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	mockRepo.On("GetTestCaseByID", "999").Return(entities.TestCase{}, assert.AnError)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-cases/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTestCase_Success(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	req := entities.CreateTestCaseRequest{Title: "Updated Case"}
	mockRepo.On("UpdateTestCase", "tc-1", req).Return(entities.TestCase{
		ID: "tc-1", Title: "Updated Case",
	}, nil)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/test-cases/tc-1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp entities.TestCase
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "Updated Case", resp.Title)

	mockRepo.AssertExpectations(t)
}

func TestUpdateTestCase_Error(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	req := entities.CreateTestCaseRequest{Title: "Updated Case"}
	mockRepo.On("UpdateTestCase", "tc-1", req).Return(entities.TestCase{}, assert.AnError)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/test-cases/tc-1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTestCase_Success(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	mockRepo.On("DeleteTestCase", "tc-1").Return(nil)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/test-cases/tc-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTestCase_Error(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	mockRepo.On("DeleteTestCase", "tc-1").Return(assert.AnError)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/test-cases/tc-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetTestCaseHistory_Success(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	mockRepo.On("GetTestCaseHistory", "tc-1").Return([]entities.TestCaseHistory{
		{ID: "h-1", TestCaseID: "tc-1"},
	}, nil)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-cases/tc-1/history", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body []entities.TestCaseHistory
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Len(t, body, 1)

	mockRepo.AssertExpectations(t)
}

func TestGetTestCaseHistory_Error(t *testing.T) {
	mockRepo := new(mocks.TestCaseRepository)
	mockRepo.On("GetTestCaseHistory", "tc-1").Return(nil, assert.AnError)

	r := setupTestCaseRouter(endpoints.NewTestCaseHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-cases/tc-1/history", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}
