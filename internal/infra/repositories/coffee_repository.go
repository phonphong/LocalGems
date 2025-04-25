package repositories

import models "localgems/internal/core/entity"

type CoffeeRepository interface {
	FindAll() ([]models.Coffee, error)
	FindByID(id int) (*models.Coffee, error)
	Create(coffee *models.Coffee) (int, error)
	Update(id int, coffee *models.Coffee) error
	Delete(id int) error
	Search(query string) ([]models.Coffee, error)
}
