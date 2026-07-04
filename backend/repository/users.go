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
