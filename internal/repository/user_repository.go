package repository

import "localgems/internal/model"

// UserRepository định nghĩa các phương thức thao tác với bảng User
type UserRepository interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
}