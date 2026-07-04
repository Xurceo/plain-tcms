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
	"github.com/xurceo/plain-tcms/db"
)

type GenericRepository[T any] struct{}

func (r *GenericRepository[T]) GetAll(model *[]T) error {
	return db.DB.Find(model).Error
}

func (r *GenericRepository[T]) GetByID(id string, model *T) error {
	return db.DB.First(model, "id = ?", id).Error
}

func (r *GenericRepository[T]) Delete(id string, model *T) error {
	return db.DB.Delete(model, "id = ?", id).Error
}
