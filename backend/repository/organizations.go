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
