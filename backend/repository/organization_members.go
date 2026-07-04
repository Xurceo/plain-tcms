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
