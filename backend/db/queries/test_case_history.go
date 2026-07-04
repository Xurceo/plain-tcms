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

const TestCaseHistoryColumns = "id, test_case_id, snapshot, changed_by, changed_at"

const (
	GetHistoryByTestCase  = "SELECT " + TestCaseHistoryColumns + " FROM test_case_history WHERE test_case_id = $1 ORDER BY changed_at DESC"
	InsertTestCaseHistory = "INSERT INTO test_case_history (test_case_id, snapshot, changed_by) VALUES ($1, $2, $3) RETURNING " + TestCaseHistoryColumns
)
