package repository

import (
	"github.com/xurceo/plain-tcms/entities"
	"gorm.io/gorm"
)

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepo(db *gorm.DB) OrganizationRepository {
	return &organizationRepository{db: db}
}

func (r *organizationRepository) GetAllOrganizations() ([]entities.Organization, error) {
	var orgs []entities.Organization
	err := r.db.Find(&orgs).Error
	return orgs, err
}

func (r *organizationRepository) GetOrganizationByID(id string) (entities.Organization, error) {
	var org entities.Organization
	err := r.db.First(&org, "id = ?", id).Error
	return org, err
}

func (r *organizationRepository) CreateOrganization(req entities.CreateOrganizationRequest) (entities.Organization, error) {
	org := entities.Organization{Name: req.Name}
	err := r.db.Create(&org).Error
	return org, err
}

func (r *organizationRepository) DeleteOrganization(id string) error {
	return r.db.Delete(&entities.Organization{}, "id = ?", id).Error
}

func (r *organizationRepository) GetMembers(orgID string) ([]entities.OrganizationMember, error) {
	var members []entities.OrganizationMember
	err := r.db.Where("org_id = ?", orgID).Find(&members).Error
	return members, err
}

func (r *organizationRepository) AddMember(orgID string, req entities.CreateOrganizationMemberRequest) (entities.OrganizationMember, error) {
	member := entities.OrganizationMember{
		OrgID:  orgID,
		UserID: req.UserID,
		Role:   req.Role,
	}
	err := r.db.Create(&member).Error
	return member, err
}

func (r *organizationRepository) RemoveMember(orgID string, userID string) error {
	return r.db.Where("org_id = ? AND user_id = ?", orgID, userID).Delete(&entities.OrganizationMember{}).Error
}
