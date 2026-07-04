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

package queries

const OrganizationColumns = "id, name, created_at"

const (
	GetAllOrganizations = "SELECT " + OrganizationColumns + " FROM organizations"
	GetOrganizationByID = "SELECT " + OrganizationColumns + " FROM organizations WHERE id = $1"
	InsertOrganization  = "INSERT INTO organizations (name) VALUES ($1) RETURNING " + OrganizationColumns
	DeleteOrganization  = "DELETE FROM organizations WHERE id = $1"
)
