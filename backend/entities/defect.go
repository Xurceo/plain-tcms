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

type Defect struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	ResultID     *string   `gorm:"column:result_id;type:uuid" json:"result_id"`
	ExternalLink *string   `gorm:"column:external_link" json:"external_link"`
	Title        string    `gorm:"not null" json:"title"`
	Severity     *string   `json:"severity"`
	Status       string    `gorm:"default:open" json:"status"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateDefectRequest struct {
	ResultID     *string `json:"result_id"`
	ExternalLink *string `json:"external_link"`
	Title        string  `json:"title"`
	Severity     *string `json:"severity"`
}

func (Defect) TableName() string { return "defects" }
