package controllers

import (
	"LifeGuard/src/mpu6050/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllMpu6050Controller struct {
	getAllUseCase *application.GetAllMpu6050UseCase
}

func NewGetAllMpu6050Controller(mpu6050Service *application.GetAllMpu6050UseCase) *GetAllMpu6050Controller {
	return &GetAllMpu6050Controller{getAllUseCase: mpu6050Service}
}

func (c *GetAllMpu6050Controller) GetMpu6050Data(ctx *gin.Context) {
	data, err := c.getAllUseCase.GetAllMpu6050()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}
