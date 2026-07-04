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

const UserColumns = "id, email, created_at"

const (
	GetAllUsers    = "SELECT " + UserColumns + " FROM users"
	GetUserByID    = "SELECT " + UserColumns + " FROM users WHERE id = $1"
	GetUserByEmail = "SELECT id, email, password_hash, created_at FROM users WHERE email = $1"
	InsertUser     = "INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING " + UserColumns
	DeleteUser     = "DELETE FROM users WHERE id = $1"
)
