package application

import (
	"LifeGuard/src/mpu6050/domain/entities"
	"LifeGuard/src/mpu6050/domain/repository"
)

type GetAllMpu6050UseCase struct{
	mpu6050Repo repository.Mpu6050Repository
}

func NewGetAllMpu6050UseCase (mpu6050 repository.Mpu6050Repository) *GetAllMpu6050UseCase {
	return &GetAllMpu6050UseCase{
		mpu6050Repo: mpu6050,
	}
}

func (s *GetAllMpu6050UseCase) GetAllMpu6050() ([]*entities.Mpu6050, error) {
	return s.mpu6050Repo.GetAll()		
}