package model

import (
	"time"
)

// Place đại diện cho một địa điểm (như quán cà phê)
type Place struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Address     string    `json:"address" gorm:"not null"`
	Latitude    float64   `json:"latitude" gorm:"not null"`
	Longitude   float64   `json:"longitude" gorm:"not null"`
	PriceRange  string    `json:"price_range"`
	Type        string    `json:"type"`  // ví dụ: "Quán cà phê", "Nhà hàng"
	Photos      []Photo   `json:"photos" gorm:"foreignKey:PlaceID"`
	Reviews     []Review  `json:"reviews,omitempty" gorm:"foreignKey:PlaceID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Photo đại diện cho ảnh của một địa điểm
type Photo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	PlaceID   uint      `json:"place_id" gorm:"not null"`
	URL       string    `json:"url" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}

// PlaceResponse là cấu trúc dữ liệu trả về cho client
type PlaceResponse struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Address     string        `json:"address"`
	Latitude    float64       `json:"latitude"`
	Longitude   float64       `json:"longitude"`
	PriceRange  string        `json:"price_range"`
	Type        string        `json:"type"`
	Photos      []Photo       `json:"photos"`
	Rating      float64       `json:"rating"` // Tính trung bình từ reviews
	ReviewCount int           `json:"review_count"`
	Reviews     []ReviewBrief `json:"reviews,omitempty"`
}

// ToResponse chuyển đổi Place thành PlaceResponse
func (p *Place) ToResponse(includeReviews bool) PlaceResponse {
	response := PlaceResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Address:     p.Address,
		Latitude:    p.Latitude,
		Longitude:   p.Longitude,
		PriceRange:  p.PriceRange,
		Type:        p.Type,
		Photos:      p.Photos,
		ReviewCount: len(p.Reviews),
	}

	// Tính trung bình rating
	var totalRating float64
	for _, review := range p.Reviews {
		totalRating += float64(review.Rating)
	}
	if len(p.Reviews) > 0 {
		response.Rating = totalRating / float64(len(p.Reviews))
	}

	// Thêm reviews nếu cần
	if includeReviews && len(p.Reviews) > 0 {
		response.Reviews = make([]ReviewBrief, len(p.Reviews))
		for i, review := range p.Reviews {
			response.Reviews[i] = review.ToBrief()
		}
	}

	return response
}

// SearchRequest là cấu trúc dữ liệu đầu vào cho tìm kiếm địa điểm
type SearchRequest struct {
	Query    string  `form:"q"`
	MinRating float64 `form:"rating"`
	Type     string  `form:"type"`
	Limit    int     `form:"limit,default=10"`
	Offset   int     `form:"offset,default=0"`
}

// CreatePlaceRequest là cấu trúc dữ liệu đầu vào cho tạo địa điểm mới
type CreatePlaceRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Address     string   `json:"address" binding:"required"`
	Latitude    float64  `json:"latitude" binding:"required"`
	Longitude   float64  `json:"longitude" binding:"required"`
	PriceRange  string   `json:"price_range"`
	Type        string   `json:"type" binding:"required"`
	Photos      []string `json:"photos"`
}