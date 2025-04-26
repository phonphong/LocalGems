package repositories

import "localgems/internal/core/entity"

type CoffeeRepository interface {
	FindAll() ([]entity.Coffee, error)
	FindByID(id int) (*entity.Coffee, error)
	Create(coffee *entity.Coffee) (int, error)
	Update(id int, coffee *entity.Coffee) error
	Delete(id int) error
	Search(query string) ([]entity.Coffee, error)
}
