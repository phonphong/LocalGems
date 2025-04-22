package repository

import (
	"localgems/internal/model"
	"localgems/internal/repository"
	"gorm.io/gorm"
)

// reviewRepo là implementation của ReviewRepository interface
type reviewRepo struct {
	db *gorm.DB
}

// NewReviewRepository tạo instance mới của ReviewRepository
func NewReviewRepository(db *gorm.DB) repository.ReviewRepository {
	return &reviewRepo{db: db}
}

// Create tạo đánh giá mới
func (r *reviewRepo) Create(review *model.Review) error {
	return r.db.Create(review).Error
}

// GetByID tìm đánh giá theo ID
func (r *reviewRepo) GetByID(id uint) (*model.Review, error) {
	var review model.Review
	err := r.db.Preload("User").First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// GetByPlaceID lấy danh sách đánh giá theo địa điểm
func (r *reviewRepo) GetByPlaceID(placeID uint, limit, offset int) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Preload("User").
		Where("place_id = ?", placeID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&reviews).Error
	return reviews, err
}

// GetByUserID lấy danh sách đánh giá theo người dùng
func (r *reviewRepo) GetByUserID(userID uint, limit, offset int) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Preload("User").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&reviews).Error
	return reviews, err
}

// Update cập nhật đánh giá
func (r *reviewRepo) Update(review *model.Review) error {
	return r.db.Save(review).Error
}

// Delete xóa đánh giá
func (r *reviewRepo) Delete(id uint) error {
	return r.db.Delete(&model.Review{}, id).Error
}