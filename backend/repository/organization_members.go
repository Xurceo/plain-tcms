package repository

import (
	"github.com/xurceo/plain-tcms/entities"
	"gorm.io/gorm"
)

type organizationMemberRepository struct {
	db *gorm.DB
}

func NewOrganizationMemberRepo(db *gorm.DB) *organizationMemberRepository {
	return &organizationMemberRepository{db: db}
}

func (r *organizationMemberRepository) GetMembersByOrg(orgID string) ([]entities.OrganizationMember, error) {
	var members []entities.OrganizationMember
	err := r.db.Where("org_id = ?", orgID).Find(&members).Error
	return members, err
}

func (r *organizationMemberRepository) GetMembersByUser(userID string) ([]entities.OrganizationMember, error) {
	var members []entities.OrganizationMember
	err := r.db.Where("user_id = ?", userID).Find(&members).Error
	return members, err
}

func (r *organizationMemberRepository) AddMember(orgID string, req entities.CreateOrganizationMemberRequest) (entities.OrganizationMember, error) {
	member := entities.OrganizationMember{
		OrgID:  orgID,
		UserID: req.UserID,
		Role:   req.Role,
	}
	err := r.db.Create(&member).Error
	return member, err
}

func (r *organizationMemberRepository) RemoveMember(orgID string, userID string) error {
	return r.db.Where("org_id = ? AND user_id = ?", orgID, userID).Delete(&entities.OrganizationMember{}).Error
}
