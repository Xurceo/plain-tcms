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
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/xurceo/plain-tcms/entities"
	"gorm.io/gorm"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func hashPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return hex.EncodeToString(h[:])
}

func (r *userRepository) Register(email, password string) (entities.User, error) {
	user := entities.User{
		Email:        email,
		PasswordHash: hashPassword(password),
	}
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) Login(email, password string) (entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	if user.PasswordHash != hashPassword(password) {
		return user, ErrInvalidCredentials
	}
	return user, nil
}
