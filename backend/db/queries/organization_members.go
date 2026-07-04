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

const OrgMemberColumns = "org_id, user_id, role"

const (
	GetMembersByOrg     = "SELECT " + OrgMemberColumns + " FROM organization_members WHERE org_id = $1"
	GetMembersByUser    = "SELECT " + OrgMemberColumns + " FROM organization_members WHERE user_id = $1"
	InsertOrgMember     = "INSERT INTO organization_members (org_id, user_id, role) VALUES ($1, $2, $3) RETURNING " + OrgMemberColumns
	DeleteOrgMember     = "DELETE FROM organization_members WHERE org_id = $1 AND user_id = $2"
	UpdateOrgMemberRole = "UPDATE organization_members SET role = $1 WHERE org_id = $2 AND user_id = $3"
)
