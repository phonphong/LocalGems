package impl

import (
	"errors"
	"localgems/internal/model"
	"localgems/internal/repository"
	"localgems/
	

// placeService là implementation của PlaceService interface
type placeService struct {
	placeRepo repository.PlaceRepository
}

// NewPlaceService tạo instance mới của PlaceService
func NewPlaceService(placeRepo repository.PlaceRepository) service.PlaceService {
	return &placeService{
		placeRepo: placeRepo,
	}
}

// CreatePlace xử lý tạo địa điểm mới
func (s *placeService) CreatePlace(req model.CreatePlaceRequest) (*model.PlaceResponse, error) {
	// Tạo đối tượng Place mới
	newPlace := &model.Place{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		PriceRange:  req.PriceRange,
		Type:        req.Type,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Lưu địa điểm vào database
	if err := s.placeRepo.Create(newPlace); err != nil {
		return nil, errors.New("không thể tạo địa điểm")
	}

	// Thêm ảnh nếu có
	for _, photoURL := range req.Photos {
		photo := &model.Photo{
			PlaceID:   newPlace.ID,
			URL:       photoURL,
			CreatedAt: time.Now(),
		}
		if err := s.placeRepo.AddPhoto(photo); err != nil {
			// Log lỗi nhưng không ngừng xử lý
			continue
		}
		newPlace.Photos = append(newPlace.Photos, *photo)
	}

	// Chuyển đổi sang response và trả về
	response := newPlace.ToResponse(false)
	return &response, nil
}

// GetPlaceByID lấy thông tin địa điểm theo ID
func (s *placeService) GetPlaceByID(id uint, includeReviews bool) (*model.PlaceResponse, error) {
	// Lấy địa điểm từ database
	place, err := s.placeRepo.GetByID(id, includeReviews)
	if err != nil {
		return nil, errors.New("không tìm thấy địa điểm")
	}

	// Chuyển đổi sang response và trả về
	response := place.ToResponse(includeReviews)
	return &response, nil
}

// GetAllPlaces lấy danh sách địa điểm với phân trang
func (s *placeService) GetAllPlaces(limit, offset int) ([]model.PlaceResponse, error) {
	// Lấy danh sách từ database
	places, err := s.placeRepo.GetAll(limit, offset)
	if err != nil {
		return nil, errors.New("lỗi khi lấy danh sách địa điểm")
	}

	// Chuyển đổi sang response
	var responses []model.PlaceResponse
	for _, place := range places {
		placeCopy := place // Tạo bản sao để tránh vấn đề với địa chỉ bộ nhớ trong vòng lặp
		responses = append(responses, placeCopy.ToResponse(false))
	}

	return responses, nil
}

// SearchPlaces tìm kiếm địa điểm theo các tiêu chí
func (s *placeService) SearchPlaces(req model.SearchRequest) ([]model.PlaceResponse, error) {
	// Thiết lập giá trị mặc định nếu không có
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	// Tìm kiếm từ database
	places, err := s.placeRepo.Search(req.Query, req.MinRating, req.Type, limit, req.Offset)
	if err != nil {
		return nil, errors.New("lỗi khi tìm kiếm địa điểm")
	}

	// Chuyển đổi sang response
	var responses []model.PlaceResponse
	for _, place := range places {
		placeCopy := place
		responses = append(responses, placeCopy.ToResponse(false))
	}

	return responses, nil
}

// AddPhotoToPlace thêm ảnh cho địa điểm
func (s *placeService) AddPhotoToPlace(placeID uint, photoURL string) error {
	// Kiểm tra địa điểm tồn tại
	_, err := s.placeRepo.GetByID(placeID, false)
	if err != nil {
		return errors.New("không tìm thấy địa điểm")
	}

	// Thêm ảnh
	photo := &model.Photo{
		PlaceID:   placeID,
		URL:       photoURL,
		CreatedAt: time.Now(),
	}

	if err := s.placeRepo.AddPhoto(photo); err != nil {
		return errors.New("không thể thêm ảnh")
	}

	return nil
}
