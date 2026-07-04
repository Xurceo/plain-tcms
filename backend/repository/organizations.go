package repository

import (
	"github.com/xurceo/plain-tcms/db"
	"github.com/xurceo/plain-tcms/entities"
)

type OrganizationRepo struct {
	*GenericRepository[entities.Organization]
}

func GetAllOrganizations() ([]entities.Organization, error) {
	var orgs []entities.Organization
	err := db.DB.Find(&orgs).Error
	return orgs, err
}

func GetOrganizationByID(id string) (entities.Organization, error) {
	var org entities.Organization
	err := db.DB.First(&org, "id = ?", id).Error
	return org, err
}

func CreateOrganization(req entities.CreateOrganizationRequest) (entities.Organization, error) {
	org := entities.Organization{Name: req.Name}
	err := db.DB.Create(&org).Error
	return org, err
}

func DeleteOrganization(id string) error {
	return db.DB.Delete(&entities.Organization{}, "id = ?", id).Error
}
