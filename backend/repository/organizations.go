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
