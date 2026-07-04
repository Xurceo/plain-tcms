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

package repository

import (
	"github.com/xurceo/plain-tcms/entities"
	"gorm.io/gorm"
)

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) GetAllProjects() ([]entities.Project, error) {
	var projects []entities.Project
	err := r.db.Find(&projects).Error
	return projects, err
}

func (r *projectRepository) GetProjectByID(id string) (entities.Project, error) {
	var project entities.Project
	err := r.db.First(&project, "id = ?", id).Error
	return project, err
}

func (r *projectRepository) GetProjectsByOrgID(orgID string) ([]entities.Project, error) {
	var projects []entities.Project
	err := r.db.Where("org_id = ?", orgID).Find(&projects).Error
	return projects, err
}

func (r *projectRepository) CreateProject(req entities.CreateProjectRequest) (entities.Project, error) {
	project := entities.Project{
		OrgID:       req.OrgID,
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
	}
	err := r.db.Create(&project).Error
	return project, err
}

func (r *projectRepository) UpdateProject(id string, req entities.CreateProjectRequest) (entities.Project, error) {
	var project entities.Project
	if err := r.db.First(&project, "id = ?", id).Error; err != nil {
		return project, err
	}
	project.Name = req.Name
	project.Description = req.Description
	project.CreatedBy = req.CreatedBy
	err := r.db.Save(&project).Error
	return project, err
}

func (r *projectRepository) DeleteProject(id string) error {
	return r.db.Delete(&entities.Project{}, "id = ?", id).Error
}
