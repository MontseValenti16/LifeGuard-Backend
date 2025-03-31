package application

import (
	"LifeGuard/src/mpu6050/domain/entities"
	"LifeGuard/src/mpu6050/domain/repository"
)

type CreateMpu6050UseCase struct {
	mpu6050Repo repository.Mpu6050Repository
}

func NewCreateMpu6050UseCase(mpu6050repo repository.Mpu6050Repository) *CreateMpu6050UseCase {
	return &CreateMpu6050UseCase{
		mpu6050Repo: mpu6050repo,
	}
}

func (s *CreateMpu6050UseCase) Run(data *entities.Mpu6050) error {
	return s.mpu6050Repo.Save(data)
}
