package application

import (
	"LifeGuard/src/ds18b20/domain/entities"
	"LifeGuard/src/ds18b20/domain/repository"
)

type GetAllDs18b20UseCase struct{
	ds18b20Repo repository.Ds18b20Repository
}

func NewGetAllDs18b20UseCase (ds18b20 repository.Ds18b20Repository) *GetAllDs18b20UseCase {
	return &GetAllDs18b20UseCase{
		ds18b20Repo: ds18b20,
	}
}

func (s *GetAllDs18b20UseCase) GetAllDs18b20() ([]*entities.Ds18b20, error) {
	return s.ds18b20Repo.GetAll()		
}