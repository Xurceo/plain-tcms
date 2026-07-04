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

func setupDefectRouter(h *endpoints.DefectHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/defects/:id", h.GetDefectByID)
	r.PUT("/defects/:id", h.UpdateDefect)
	r.DELETE("/defects/:id", h.DeleteDefect)
	return r
}

func TestGetDefectByID_Success(t *testing.T) {
	mockRepo := new(mocks.DefectRepository)
	mockRepo.On("GetDefectByID", "d-1").Return(entities.Defect{
		ID: "d-1", Title: "Bug One",
	}, nil)

	r := setupDefectRouter(endpoints.NewDefectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/defects/d-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body entities.Defect
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Equal(t, "Bug One", body.Title)

	mockRepo.AssertExpectations(t)
}

func TestGetDefectByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.DefectRepository)
	mockRepo.On("GetDefectByID", "999").Return(entities.Defect{}, assert.AnError)

	r := setupDefectRouter(endpoints.NewDefectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/defects/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestUpdateDefect_Success(t *testing.T) {
	mockRepo := new(mocks.DefectRepository)
	req := entities.CreateDefectRequest{Title: "Updated Bug"}
	mockRepo.On("UpdateDefect", "d-1", req).Return(entities.Defect{
		ID: "d-1", Title: "Updated Bug",
	}, nil)

	r := setupDefectRouter(endpoints.NewDefectHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/defects/d-1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp entities.Defect
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "Updated Bug", resp.Title)

	mockRepo.AssertExpectations(t)
}

func TestUpdateDefect_Error(t *testing.T) {
	mockRepo := new(mocks.DefectRepository)
	req := entities.CreateDefectRequest{Title: "Updated Bug"}
	mockRepo.On("UpdateDefect", "d-1", req).Return(entities.Defect{}, assert.AnError)

	r := setupDefectRouter(endpoints.NewDefectHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/defects/d-1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteDefect_Success(t *testing.T) {
	mockRepo := new(mocks.DefectRepository)
	mockRepo.On("DeleteDefect", "d-1").Return(nil)

	r := setupDefectRouter(endpoints.NewDefectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/defects/d-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteDefect_Error(t *testing.T) {
	mockRepo := new(mocks.DefectRepository)
	mockRepo.On("DeleteDefect", "d-1").Return(assert.AnError)

	r := setupDefectRouter(endpoints.NewDefectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/defects/d-1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}
