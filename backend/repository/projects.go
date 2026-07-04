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
