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

type TestResult struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	RunID      string    `gorm:"column:run_id;type:uuid;not null" json:"run_id"`
	TestCaseID string    `gorm:"column:test_case_id;type:uuid;not null" json:"test_case_id"`
	Status     string    `gorm:"default:untested" json:"status"`
	Comment    *string   `json:"comment"`
	ExecutedBy *string   `gorm:"column:executed_by;type:uuid" json:"executed_by"`
	DurationMs *int      `gorm:"column:duration_ms" json:"duration_ms"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type CreateTestResultRequest struct {
	TestCaseID string  `json:"test_case_id"`
	Status     string  `json:"status"`
	Comment    *string `json:"comment"`
	DurationMs *int    `json:"duration_ms"`
}

func (TestResult) TableName() string { return "test_results" }
