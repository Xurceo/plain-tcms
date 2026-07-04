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

func setupOrganizationRouter(h *endpoints.OrganizationHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/organizations", h.GetAllOrganizations)
	r.GET("/organizations/:id", h.GetOrganizationByID)
	r.POST("/organizations", h.CreateOrganization)
	r.DELETE("/organizations/:id", h.DeleteOrganization)
	return r
}

func TestGetAllOrganizations_Success(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	mockRepo.On("GetAllOrganizations").Return([]entities.Organization{
		{ID: "1", Name: "Org One"},
		{ID: "2", Name: "Org Two"},
	}, nil)

	r := setupOrganizationRouter(endpoints.NewOrganizationHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/organizations", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body []entities.Organization
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Len(t, body, 2)

	mockRepo.AssertExpectations(t)
}

func TestGetOrganizationByID_Success(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	mockRepo.On("GetOrganizationByID", "1").Return(entities.Organization{
		ID: "1", Name: "Org One",
	}, nil)

	r := setupOrganizationRouter(endpoints.NewOrganizationHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/organizations/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body entities.Organization
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Equal(t, "Org One", body.Name)

	mockRepo.AssertExpectations(t)
}

func TestGetOrganizationByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	mockRepo.On("GetOrganizationByID", "999").Return(entities.Organization{}, assert.AnError)

	r := setupOrganizationRouter(endpoints.NewOrganizationHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/organizations/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateOrganization_Success(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	req := entities.CreateOrganizationRequest{Name: "New Org"}
	mockRepo.On("CreateOrganization", req).Return(entities.Organization{
		ID: "3", Name: "New Org",
	}, nil)

	r := setupOrganizationRouter(endpoints.NewOrganizationHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/organizations", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entities.Organization
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "New Org", resp.Name)

	mockRepo.AssertExpectations(t)
}

func TestDeleteOrganization_Success(t *testing.T) {
	mockRepo := new(mocks.OrganizationRepository)
	mockRepo.On("DeleteOrganization", "1").Return(nil)

	r := setupOrganizationRouter(endpoints.NewOrganizationHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/organizations/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockRepo.AssertExpectations(t)
}
