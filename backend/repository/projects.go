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
	"github.com/xurceo/plain-tcms/db"
	"github.com/xurceo/plain-tcms/entities"
)

func GetAllProjects() ([]entities.Project, error) {
	var projects []entities.Project
	err := db.DB.Find(&projects).Error
	return projects, err
}

func GetProjectByID(projectID string) (entities.Project, error) {
	var project entities.Project
	err := db.DB.First(&project, "id = ?", projectID).Error
	return project, err
}

func GetProjectsByOrgID(orgID string) ([]entities.Project, error) {
	var projects []entities.Project
	err := db.DB.Where("org_id = ?", orgID).Find(&projects).Error
	return projects, err
}

func CreateProject(req entities.CreateProjectRequest) (entities.Project, error) {
	project := entities.Project{
		OrgID:       req.OrgID,
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
	}
	err := db.DB.Create(&project).Error
	return project, err
}

func DeleteProject(projectID string) error {
	return db.DB.Delete(&entities.Project{}, "id = ?", projectID).Error
}
