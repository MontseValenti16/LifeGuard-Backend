package application

import (
	"LifeGuard/src/max30102/domain/entities"
	repositories "LifeGuard/src/max30102/domain/repository"
)

type CreateMax30102UseCase struct {
    max30102Repo repositories.Max30102Repository
}

func NewMax30102UseCase ( max30102Repo repositories.Max30102Repository ) *CreateMax30102UseCase {
    return &CreateMax30102UseCase {
        max30102Repo: max30102Repo,
    }
}

func (s *CreateMax30102UseCase ) Run(data *entities.Max30102) error {
    return s.max30102Repo.Save(data)
}

