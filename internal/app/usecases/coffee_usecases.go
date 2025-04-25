package usecases

import (
	models "localgems/internal/core/entity"
	"localgems/internal/infra/repositories"
)

type CafeUsecase interface {
	GetAllCafes() ([]models.Cafe, error)
	GetCafeByID(id int) (*models.Cafe, error)
	CreateCafe(cafe *models.Cafe) (int, error)
	UpdateCafe(id int, cafe *models.Cafe) error
	DeleteCafe(id int) error
	SearchCafes(query string) ([]models.Cafe, error)
}

type cafeUsecase struct {
	cafeRepo repositories.CafeRepository
}

func NewCafeUsecase(repo repositories.CafeRepository) CafeUsecase {
	return &cafeUsecase{
		cafeRepo: repo,
	}
}

func (u *cafeUsecase) GetAllCafes() ([]models.Cafe, error) {
	return u.cafeRepo.FindAll()
}

func (u *cafeUsecase) GetCafeByID(id int) (*models.Cafe, error) {
	cafe, err := u.cafeRepo.FindByID(id)
	if err != nil {
		return nil, errors.NewNotFoundError("cafe not found")
	}
	return cafe, nil
}

func (u *cafeUsecase) CreateCafe(cafe *models.Cafe) (int, error) {
	// Validate cafe if needed
	if cafe.Name == "" {
		return 0, errors.NewValidationError("name is required")
	}
	return u.cafeRepo.Create(cafe)
}

func (u *cafeUsecase) UpdateCafe(id int, cafe *models.Cafe) error {
	// Check if cafe exists
	_, err := u.cafeRepo.FindByID(id)
	if err != nil {
		return errors.NewNotFoundError("cafe not found")
	}

	// Validate cafe if needed
	if cafe.Name == "" {
		return errors.NewValidationError("name is required")
	}

	return u.cafeRepo.Update(id, cafe)
}

func (u *cafeUsecase) DeleteCafe(id int) error {
	// Check if cafe exists
	_, err := u.cafeRepo.FindByID(id)
	if err != nil {
		return errors.NewNotFoundError("cafe not found")
	}

	return u.cafeRepo.Delete(id)
}

func (u *cafeUsecase) SearchCafes(query string) ([]models.Cafe, error) {
	if query == "" {
		return nil, errors.NewValidationError("search query is required")
	}
	return u.cafeRepo.Search(query)
}
