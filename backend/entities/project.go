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

package entities

import "time"

type Project struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	OrgID       string    `gorm:"column:org_id;type:uuid;not null" json:"org_id"`
	Name        string    `gorm:"not null" json:"name"`
	Description *string   `json:"description"`
	CreatedBy   *string   `gorm:"column:created_by;type:uuid" json:"created_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateProjectRequest struct {
	OrgID       string  `json:"org_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedBy   *string `json:"created_by"`
}

func (Project) TableName() string { return "projects" }
