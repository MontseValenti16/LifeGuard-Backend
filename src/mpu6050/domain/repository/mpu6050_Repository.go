package repository

import "LifeGuard/src/mpu6050/domain/entities"

type Mpu6050Repository interface {
	Save(data *entities.Mpu6050) error
	GetAll() ([]*entities.Mpu6050, error)
}