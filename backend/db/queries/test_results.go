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

const TestResultColumns = "id, run_id, test_case_id, status, comment, executed_by, duration_ms, created_at"

const (
	GetTestResultsByRun = "SELECT " + TestResultColumns + " FROM test_results WHERE run_id = $1"
	GetTestResultByID   = "SELECT " + TestResultColumns + " FROM test_results WHERE id = $1"
	InsertTestResult    = "INSERT INTO test_results (run_id, test_case_id, status, comment, executed_by, duration_ms) VALUES ($1, $2, $3, $4, $5, $6) RETURNING " + TestResultColumns
	DeleteTestResult    = "DELETE FROM test_results WHERE id = $1"
)
