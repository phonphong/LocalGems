package usecases

import (
	"local-gems-server/internal/core/entity"
	"local-gems-server/internal/core/errors"
	"local-gems-server/internal/infra/repositories"
)

type LocalUsecase interface {
	GetAllLocals() ([]entity.Local, error)
	GetLocalByID(id int) (*entity.Local, error)
	CreateLocal(local *entity.Local) (int, error)
	UpdateLocal(id int, local *entity.Local) error
	DeleteLocal(id int) error
	SearchLocals(query string) ([]entity.Local, error)
}

type localUsecase struct {
	localRepo repositories.LocalRepository
}

func NewLocalUsecase(repo repositories.LocalRepository) LocalUsecase {
	return &localUsecase{
		localRepo: repo,
	}
}

func (u *localUsecase) GetAllLocals() ([]entity.Local, error) {
	return u.localRepo.FindAll()
}

func (u *localUsecase) GetLocalByID(id int) (*entity.Local, error) {
	local, err := u.localRepo.FindByID(id)
	if err != nil {
		return nil, errors.NewNotFoundError("local not found")
	}
	return local, nil
}

func (u *localUsecase) CreateLocal(local *entity.Local) (int, error) {
	if local.Name == "" {
		return 0, errors.NewValidationError("name is required")
	}
	return u.localRepo.Create(local)
}

func (u *localUsecase) UpdateLocal(id int, local *entity.Local) error {
	_, err := u.localRepo.FindByID(id)
	if err != nil {
		return errors.NewNotFoundError("local not found")
	}

	if local.Name == "" {
		return errors.NewValidationError("name is required")
	}

	return u.localRepo.Update(id, local)
}

func (u *localUsecase) DeleteLocal(id int) error {
	_, err := u.localRepo.FindByID(id)
	if err != nil {
		return errors.NewNotFoundError("local not found")
	}
	return u.localRepo.Delete(id)
}

func (u *localUsecase) SearchLocals(query string) ([]entity.Local, error) {
	if query == "" {
		return nil, errors.NewValidationError("search query is required")
	}
	return u.localRepo.Search(query)
}
