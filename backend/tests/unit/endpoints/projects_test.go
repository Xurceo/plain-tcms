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

func setupProjectRouter(h *endpoints.ProjectHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/projects", h.GetProjects)
	r.GET("/projects/:id", h.GetProjectByID)
	r.GET("/organizations/:id/projects", h.GetProjectsByOrgID)
	r.POST("/organizations/:id/projects", h.CreateProject)
	r.PUT("/projects/:id", h.UpdateProject)
	r.DELETE("/projects/:id", h.DeleteProject)
	return r
}

func TestGetProjects_Success(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	mockRepo.On("GetAllProjects").Return([]entities.Project{
		{ID: "1", Name: "Proj One"},
		{ID: "2", Name: "Proj Two"},
	}, nil)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/projects", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body []entities.Project
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Len(t, body, 2)

	mockRepo.AssertExpectations(t)
}

func TestGetProjects_Error(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	mockRepo.On("GetAllProjects").Return(nil, assert.AnError)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/projects", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetProjectByID_Success(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	mockRepo.On("GetProjectByID", "1").Return(entities.Project{
		ID: "1", Name: "Proj One",
	}, nil)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/projects/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body entities.Project
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Equal(t, "Proj One", body.Name)

	mockRepo.AssertExpectations(t)
}

func TestGetProjectByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	mockRepo.On("GetProjectByID", "999").Return(entities.Project{}, assert.AnError)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/projects/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateProject_Success(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	req := entities.CreateProjectRequest{OrgID: "org-1", Name: "New Proj"}
	mockRepo.On("CreateProject", req).Return(entities.Project{
		ID: "3", Name: "New Proj",
	}, nil)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/organizations/org-1/projects", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusCreated, w.Code)

	var resp entities.Project
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "New Proj", resp.Name)

	mockRepo.AssertExpectations(t)
}

func TestGetProjectsByOrgID_Success(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	mockRepo.On("GetProjectsByOrgID", "org-1").Return([]entities.Project{
		{ID: "1", Name: "Proj One"},
	}, nil)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/organizations/org-1/projects", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var body []entities.Project
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		return
	}
	assert.Len(t, body, 1)

	mockRepo.AssertExpectations(t)
}

func TestGetProjectsByOrgID_Error(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	mockRepo.On("GetProjectsByOrgID", "bad-id").Return(nil, assert.AnError)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodGet, "/organizations/bad-id/projects", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateProject_Error(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	req := entities.CreateProjectRequest{OrgID: "org-1", Name: "New Proj"}
	mockRepo.On("CreateProject", req).Return(entities.Project{}, assert.AnError)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/organizations/org-1/projects", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateProject_BadRequest(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodPost, "/organizations/org-1/projects", bytes.NewBuffer([]byte("invalid")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateProject_Success(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	req := entities.CreateProjectRequest{Name: "Updated Proj"}
	mockRepo.On("UpdateProject", "1", req).Return(entities.Project{
		ID: "1", Name: "Updated Proj",
	}, nil)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/projects/1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp entities.Project
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		return
	}
	assert.Equal(t, "Updated Proj", resp.Name)

	mockRepo.AssertExpectations(t)
}

func TestUpdateProject_Error(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	req := entities.CreateProjectRequest{Name: "Updated Proj"}
	mockRepo.On("UpdateProject", "1", req).Return(entities.Project{}, assert.AnError)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPut, "/projects/1", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httpReq)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProject_Success(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	mockRepo.On("DeleteProject", "1").Return(nil)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/projects/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteProject_Error(t *testing.T) {
	mockRepo := new(mocks.ProjectRepository)
	mockRepo.On("DeleteProject", "1").Return(assert.AnError)

	r := setupProjectRouter(endpoints.NewProjectHandler(mockRepo))

	req := httptest.NewRequest(http.MethodDelete, "/projects/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRepo.AssertExpectations(t)
}
