package usecases

import (
	
	"localgems/internal/core/entity"
	"localgems/internal/infra/repositories"
	"localgems/internal/core/errors"
)

type CoffeeUsecase interface {
	GetAllCoffees() ([]entity.Coffee, error)
	GetCoffeeByID(id int) (*entity.Coffee, error)
	CreateCoffee(coffee *entity.Coffee) (int, error)
	UpdateCoffee(id int, coffee *entity.Coffee) error
	DeleteCoffee(id int) error
	SearchCoffees(query string) ([]entity.Coffee, error)
}

type coffeeUsecase struct {
	coffeeRepo repositories.CoffeeRepository
}

func NewCoffeeUsecase(repo repositories.CoffeeRepository) CoffeeUsecase {
	return &coffeeUsecase{
		coffeeRepo: repo,
	}
}

func (u *coffeeUsecase) GetAllCoffees() ([]entity.Coffee, error) {
	return u.coffeeRepo.FindAll()
}

func (u *coffeeUsecase) GetCoffeeByID(id int) (*entity.Coffee, error) {
	coffee, err := u.coffeeRepo.FindByID(id)
	if err != nil {
		return nil, errors.NewNotFoundError("coffee not found")
	}
	return coffee, nil
}

func (u *coffeeUsecase) CreateCoffee(coffee *entity.Coffee) (int, error) {
	if coffee.Name == "" {
		return 0, errors.NewValidationError("name is required")
	}
	return u.coffeeRepo.Create(coffee)
}

func (u *coffeeUsecase) UpdateCoffee(id int, coffee *entity.Coffee) error {
	_, err := u.coffeeRepo.FindByID(id)
	if err != nil {
		return errors.NewNotFoundError("coffee not found")
	}

	if coffee.Name == "" {
		return errors.NewValidationError("name is required")
	}

	return u.coffeeRepo.Update(id, coffee)
}

func (u *coffeeUsecase) DeleteCoffee(id int) error {
	_, err := u.coffeeRepo.FindByID(id)
	if err != nil {
		return errors.NewNotFoundError("coffee not found")
	}
	return u.coffeeRepo.Delete(id)
}

func (u *coffeeUsecase) SearchCoffees(query string) ([]entity.Coffee, error) {
	if query == "" {
		return nil, errors.NewValidationError("search query is required")
	}
	return u.coffeeRepo.Search(query)
}
