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

const (
	GetCasesByRun     = "SELECT test_case_id FROM test_run_cases WHERE run_id = $1"
	GetRunsByCase     = "SELECT run_id FROM test_run_cases WHERE test_case_id = $1"
	InsertTestRunCase = "INSERT INTO test_run_cases (run_id, test_case_id) VALUES ($1, $2)"
	DeleteTestRunCase = "DELETE FROM test_run_cases WHERE run_id = $1 AND test_case_id = $2"
	DeleteAllRunCases = "DELETE FROM test_run_cases WHERE run_id = $1"
)
