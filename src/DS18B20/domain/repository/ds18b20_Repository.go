package repository

import "LifeGuard/src/ds18b20/domain/entities"

type Ds18b20Repository interface {
	Save(data *entities.Ds18b20) error
	GetAll() ([]*entities.Ds18b20, error)
}