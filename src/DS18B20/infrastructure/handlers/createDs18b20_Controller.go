package handlers
import (
	"LifeGuard/src/ds18b20/application"
	"LifeGuard/src/ds18b20/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateDs18b20Controller struct {
	createUseCase *application.CreateDs18b20UseCase
	//publisher     *broker.RabbitMQPublisher
}

func NewCreateDs18b20Controller(ds18b20Service *application.CreateDs18b20UseCase /* publisher  *broker.RabbitMQPublisher*/) *CreateDs18b20Controller {
	return &CreateDs18b20Controller{createUseCase: ds18b20Service /* publisher: publisher*/}
}

func (c *CreateDs18b20Controller) SaveDs18b20Data(ctx *gin.Context) {
	var data *entities.Ds18b20
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