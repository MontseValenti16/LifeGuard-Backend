package controllers

import (
	"LifeGuard/src/ds18b20/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllDs18b20Controller struct {
	getAllUseCase *application.GetAllDs18b20UseCase
}

func NewGetAllDs18b20Controller(ds18b20Service *application.GetAllDs18b20UseCase) *GetAllDs18b20Controller {
	return &GetAllDs18b20Controller{getAllUseCase: ds18b20Service}
}

func (c *GetAllDs18b20Controller) GetDs18b20Data(ctx *gin.Context) {
	data, err := c.getAllUseCase.GetAllDs18b20()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}
