package repository

import (
	"localgems/internal/repository"
	"localgems/internal/model"
	"gorm.io/gorm"
)

// userRepo là implementation của UserRepository interface
type userRepo struct {
	db *gorm.DB
}

// NewUserRepository tạo instance mới của UserRepository
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepo{db: db}
}

// Create tạo người dùng mới
func (r *userRepo) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// GetByID tìm người dùng theo ID
func (r *userRepo) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail tìm người dùng theo email
func (r *userRepo) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update cập nhật thông tin người dùng
func (r *userRepo) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete xóa người dùng
func (r *userRepo) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
