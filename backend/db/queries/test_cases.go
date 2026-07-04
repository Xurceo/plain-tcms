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

const TestCaseColumns = "id, project_id, suite_id, title, description, preconditions, steps, expected, status, priority, type, tags, created_by, created_at, updated_at"

const (
	GetTestCasesByProject = "SELECT " + TestCaseColumns + " FROM test_cases WHERE project_id = $1"
	GetTestCasesBySuite   = "SELECT " + TestCaseColumns + " FROM test_cases WHERE suite_id = $1"
	GetTestCaseByID       = "SELECT " + TestCaseColumns + " FROM test_cases WHERE id = $1"
	InsertTestCase        = "INSERT INTO test_cases (project_id, suite_id, title, description, preconditions, expected, priority, type, tags, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING " + TestCaseColumns
	DeleteTestCase        = "DELETE FROM test_cases WHERE id = $1"
)
