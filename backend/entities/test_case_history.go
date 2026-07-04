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

type TestCaseHistory struct {
	ID         string      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	TestCaseID string      `gorm:"column:test_case_id;type:uuid;not null" json:"test_case_id"`
	Snapshot   interface{} `gorm:"type:jsonb;serializer:json;not null" json:"snapshot"`
	ChangedBy  *string     `gorm:"column:changed_by;type:uuid" json:"changed_by"`
	ChangedAt  time.Time   `gorm:"autoCreateTime" json:"changed_at"`
}

func (TestCaseHistory) TableName() string { return "test_case_history" }
