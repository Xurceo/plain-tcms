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

const TestSuiteColumns = "id, project_id, parent_id, name, description, created_at"

const (
	GetTestSuitesByProject = "SELECT " + TestSuiteColumns + " FROM test_suites WHERE project_id = $1"
	GetTestSuiteByID       = "SELECT " + TestSuiteColumns + " FROM test_suites WHERE id = $1"
	InsertTestSuite        = "INSERT INTO test_suites (project_id, parent_id, name, description) VALUES ($1, $2, $3, $4) RETURNING " + TestSuiteColumns
	DeleteTestSuite        = "DELETE FROM test_suites WHERE id = $1"
)
