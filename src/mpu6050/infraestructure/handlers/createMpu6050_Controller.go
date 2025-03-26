package controllers

import (
	"LifeGuard/src/mpu6050/application"
	"LifeGuard/src/mpu6050/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMpu6050Controller struct {
	createUseCase *application.CreateMpu6050UseCase
	//publisher     *broker.RabbitMQPublisher
}

func NewMpu6050Controller(mpu6050Service *application.CreateMpu6050UseCase /* publisher  *broker.RabbitMQPublisher*/) *CreateMpu6050Controller {
	return &CreateMpu6050Controller{createUseCase: mpu6050Service /* publisher: publisher*/}
}

func (c *CreateMpu6050Controller) SaveMpu6050Data(ctx *gin.Context) {
	var data *entities.Mpu6050
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.createUseCase.Run(data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Data saved"})
}