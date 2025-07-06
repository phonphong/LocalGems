package repositories

import "local-gems-server/internal/core/entity"

type LocalRepository interface {
	FindAll() ([]entity.Local, error)
	FindByID(id int) (*entity.Local, error)
	Create(coffee *entity.Local) (int, error)
	Update(id int, coffee *entity.Local) error
	Delete(id int) error
	Search(query string) ([]entity.Local, error)
}
