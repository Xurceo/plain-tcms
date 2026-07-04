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
