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

package repository

import (
	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return BaseRepository{db: db}
}

func (r *BaseRepository) FetchAll(query string, dest any, args ...any) error {
	return r.db.Raw(query, args...).Scan(dest).Error
}

func (r *BaseRepository) FetchOne(query string, dest any, args ...any) error {
	return r.db.Raw(query, args...).First(dest).Error
}

func (r *BaseRepository) Insert(query string, dest any, args ...any) error {
	return r.db.Raw(query, args...).Scan(dest).Error
}

func (r *BaseRepository) Delete(query string, args ...any) error {
	return r.db.Exec(query, args...).Error
}
