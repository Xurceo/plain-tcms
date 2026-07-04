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

func setupTestResultRouter(h *endpoints.TestResultHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/test-results/:id/attachments", h.GetAttachmentsByResult)
	r.POST("/test-results/:id/attachments", h.AddAttachmentToResult)
	return r
}

func TestGetAttachmentsByResult_Success(t *testing.T) {
	mockRepo := new(mocks.TestResultRepository)
	fileType := "image/png"
	mockRepo.On("GetAttachmentsByResult", "result-1").Return([]entities.ResultAttachment{
		{ID: "att-1", ResultID: "result-1", FileURL: "https://example.com/img.png", FileType: &fileType},
	}, nil)

	r := setupTestResultRouter(endpoints.NewTestResultHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-results/result-1/attachments", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body []entities.ResultAttachment
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Len(t, body, 1)

	mockRepo.AssertExpectations(t)
}

func TestGetAttachmentsByResult_Error(t *testing.T) {
	mockRepo := new(mocks.TestResultRepository)
	mockRepo.On("GetAttachmentsByResult", "bad-id").Return(nil, assert.AnError)

	r := setupTestResultRouter(endpoints.NewTestResultHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/test-results/bad-id/attachments", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestAddAttachmentToResult_Success(t *testing.T) {
	mockRepo := new(mocks.TestResultRepository)
	fileType := "image/png"
	req := entities.CreateResultAttachmentRequest{
		FileURL:  "https://example.com/img.png",
		FileType: &fileType,
	}
	mockRepo.On("AddAttachmentToResult", "result-1", req).Return(entities.ResultAttachment{
		ID:       "att-1",
		ResultID: "result-1",
		FileURL:  "https://example.com/img.png",
		FileType: &fileType,
	}, nil)

	r := setupTestResultRouter(endpoints.NewTestResultHandler(mockRepo))

	payload, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/test-results/result-1/attachments", bytes.NewBuffer(payload))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entities.ResultAttachment
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "https://example.com/img.png", resp.FileURL)

	mockRepo.AssertExpectations(t)
}

func TestAddAttachmentToResult_BadRequest(t *testing.T) {
	mockRepo := new(mocks.TestResultRepository)

	r := setupTestResultRouter(endpoints.NewTestResultHandler(mockRepo))

	req := httptest.NewRequest(http.MethodPost, "/test-results/result-1/attachments", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAddAttachmentToResult_Error(t *testing.T) {
	mockRepo := new(mocks.TestResultRepository)
	req := entities.CreateResultAttachmentRequest{
		FileURL: "https://example.com/img.png",
	}
	mockRepo.On("AddAttachmentToResult", "result-1", req).Return(entities.ResultAttachment{}, assert.AnError)

	r := setupTestResultRouter(endpoints.NewTestResultHandler(mockRepo))

	payload, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/test-results/result-1/attachments", bytes.NewBuffer(payload))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}
