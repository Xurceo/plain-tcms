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

const TestRunColumns = "id, project_id, plan_id, name, status, environment, build_version, created_by, created_at, completed_at"

const (
	GetTestRunsByProject = "SELECT " + TestRunColumns + " FROM test_runs WHERE project_id = $1"
	GetTestRunByID       = "SELECT " + TestRunColumns + " FROM test_runs WHERE id = $1"
	InsertTestRun        = "INSERT INTO test_runs (project_id, plan_id, name, environment, build_version, created_by) VALUES ($1, $2, $3, $4, $5, $6) RETURNING " + TestRunColumns
	DeleteTestRun        = "DELETE FROM test_runs WHERE id = $1"
)
