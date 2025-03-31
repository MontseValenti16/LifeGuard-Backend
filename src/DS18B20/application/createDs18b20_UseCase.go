package application

import (
	"LifeGuard/src/ds18b20/domain/entities"
	"LifeGuard/src/ds18b20/domain/repository"
)

type CreateDs18b20UseCase struct {
	ds18b20Repo repository.Ds18b20Repository
}

func NewCreateds18b20UseCase(ds18b20 repository.Ds18b20Repository) *CreateDs18b20UseCase {
	return &CreateDs18b20UseCase{
		ds18b20Repo: ds18b20,
	}
}

func (s *CreateDs18b20UseCase) Run(data *entities.Ds18b20) error {
	return s.ds18b20Repo.Save(data)
}
