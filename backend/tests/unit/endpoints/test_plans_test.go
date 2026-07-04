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

func setupTestPlanRouter(h *endpoints.TestPlanHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/projects/:id/test-plans", h.GetTestPlansByProject)
	r.POST("/projects/:id/test-plans", h.CreateTestPlan)
	r.GET("/test-plans/:id", h.GetTestPlanByID)
	r.PUT("/test-plans/:id", h.UpdateTestPlan)
	r.DELETE("/test-plans/:id", h.DeleteTestPlan)
	return r
}

func TestGetTestPlansByProject_Success(t *testing.T) {
	mockRepo := new(mocks.TestPlanRepository)
	mockRepo.On("GetTestPlansByProject", "proj-1").Return([]entities.TestPlan{
		{ID: "tp-1", Name: "Plan One"},
		{ID: "tp-2", Name: "Plan Two"},
	}, nil)

	r := setupTestPlanRouter(endpoints.NewTestPlanHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/projects/proj-1/test-plans", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body []entities.TestPlan
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Len(t, body, 2)

	mockRepo.AssertExpectations(t)
}

func TestGetTestPlansByProject_Error(t *testing.T) {
	mockRepo := new(mocks.TestPlanRepository)
	mockRepo.On("GetTestPlansByProject", "bad-id").Return(nil, assert.AnError)

	r := setupTestPlanRouter(endpoints.NewTestPlanHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/projects/bad-id/test-plans", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateTestPlan_Success(t *testing.T) {
	mockRepo := new(mocks.TestPlanRepository)
	req := entities.CreateTestPlanRequest{Name: "New Plan"}
	mockRepo.On("CreateTestPlan", "proj-1", req).Return(entities.TestPlan{
		ID: "tp-3", Name: "New Plan",
	}, nil)

	r := setupTestPlanRouter(endpoints.NewTestPlanHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/projects/proj-1/test-plans", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entities.TestPlan
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "New Plan", resp.Name)

	mockRepo.AssertExpectations(t)
}

func TestCreateTestPlan_BadRequest(t *testing.T) {
	mockRepo := new(mocks.TestPlanRepository)

	r := setupTestPlanRouter(endpoints.NewTestPlanHandler(mockRepo))

	req := httptest.NewRequest(http.MethodPost, "/projects/proj-1/test-plans", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTestPlanByID_Success(t *testing.T) {
	mockRepo := new(mocks.TestPlanRepository)
	mockRepo.On("GetTestPlanByID", "tp-1").Return(entities.TestPlan{
		ID: "tp-1", Name: "Plan One",
	}, nil)

	r := setupTestPlanRouter(endpoints.NewTestPlanHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-plans/tp-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body entities.TestPlan
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Equal(t, "Plan One", body.Name)

	mockRepo.AssertExpectations(t)
}

func TestGetTestPlanByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.TestPlanRepository)
	mockRepo.On("GetTestPlanByID", "999").Return(entities.TestPlan{}, assert.AnError)

	r := setupTestPlanRouter(endpoints.NewTestPlanHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-plans/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTestPlan_Success(t *testing.T) {
	mockRepo := new(mocks.TestPlanRepository)
	req := entities.CreateTestPlanRequest{Name: "Updated Plan"}
	mockRepo.On("UpdateTestPlan", "tp-1", req).Return(entities.TestPlan{
		ID: "tp-1", Name: "Updated Plan",
	}, nil)

	r := setupTestPlanRouter(endpoints.NewTestPlanHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/test-plans/tp-1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp entities.TestPlan
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "Updated Plan", resp.Name)

	mockRepo.AssertExpectations(t)
}

func TestUpdateTestPlan_Error(t *testing.T) {
	mockRepo := new(mocks.TestPlanRepository)
	req := entities.CreateTestPlanRequest{Name: "Updated Plan"}
	mockRepo.On("UpdateTestPlan", "tp-1", req).Return(entities.TestPlan{}, assert.AnError)

	r := setupTestPlanRouter(endpoints.NewTestPlanHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/test-plans/tp-1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTestPlan_Success(t *testing.T) {
	mockRepo := new(mocks.TestPlanRepository)
	mockRepo.On("DeleteTestPlan", "tp-1").Return(nil)

	r := setupTestPlanRouter(endpoints.NewTestPlanHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/test-plans/tp-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTestPlan_Error(t *testing.T) {
	mockRepo := new(mocks.TestPlanRepository)
	mockRepo.On("DeleteTestPlan", "tp-1").Return(assert.AnError)

	r := setupTestPlanRouter(endpoints.NewTestPlanHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/test-plans/tp-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}
