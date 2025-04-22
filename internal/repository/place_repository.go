package repository
import (
	"localgems/internal/model"
)

// PlaceRepository định nghĩa các phương thức thao tác với bảng Place
type PlaceRepository interface {
	Create(place *model.Place) error
	GetByID(id uint, preloadReviews bool) (*model.Place, error)
	GetAll(limit, offset int) ([]model.Place, error)
	Update(place *model.Place) error
	Delete(id uint) error
	Search(query string, minRating float64, placeType string, limit, offset int) ([]model.Place, error)
		AddPhoto(photo *model.Photo) error
	}