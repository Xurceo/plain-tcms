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

const TestPlanColumns = "id, project_id, name, description, created_at"

const (
	GetTestPlansByProject = "SELECT " + TestPlanColumns + " FROM test_plans WHERE project_id = $1"
	GetTestPlanByID       = "SELECT " + TestPlanColumns + " FROM test_plans WHERE id = $1"
	InsertTestPlan        = "INSERT INTO test_plans (project_id, name, description) VALUES ($1, $2, $3) RETURNING " + TestPlanColumns
	DeleteTestPlan        = "DELETE FROM test_plans WHERE id = $1"
)
