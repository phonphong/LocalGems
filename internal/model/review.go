package model

import (
	"time"
)

// Review đại diện cho đánh giá của người dùng về một địa điểm
type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	PlaceID   uint      `json:"place_id" gorm:"not null"`
	Rating    int       `json:"rating" gorm:"not null"` // 1-5
	Comment   string    `json:"comment"`
	Photos    []string  `json:"photos" gorm:"type:text[]"` // Mảng URLs
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ReviewBrief là phiên bản tóm tắt của Review để hiển thị trong danh sách
type ReviewBrief struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	UserName  string    `json:"user_name"`
	UserAvatar string   `json:"user_avatar"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

// ToBrief chuyển đổi Review thành ReviewBrief
func (r *Review) ToBrief() ReviewBrief {
	return ReviewBrief{
		ID:         r.ID,
		UserID:     r.UserID,
		UserName:   r.User.Name,
		UserAvatar: r.User.Avatar,
		Rating:     r.Rating,
		Comment:    r.Comment,
		CreatedAt:  r.CreatedAt,
	}
}

// CreateReviewRequest là cấu trúc dữ liệu đầu vào cho tạo đánh giá mới
type CreateReviewRequest struct {
	PlaceID uint     `json:"place_id" binding:"required"`
	Rating  int      `json:"rating" binding:"required,min=1,max=5"`
	Comment string   `json:"comment"`
	Photos  []string `json:"photos"`
}