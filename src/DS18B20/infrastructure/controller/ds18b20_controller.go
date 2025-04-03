package controller

import (
	"LifeGuard/src/ds18b20/application"
	controllers "LifeGuard/src/ds18b20/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

type Ds18b20Controller struct {
	createProductUseCase   *controllers.CreateDs18b20Controller
	viewAllProductsUseCase *controllers.GetAllDs18b20Controller
}

func NewProductController(
	createUseCase *application.CreateDs18b20UseCase,
	viewUseCase *application.GetAllDs18b20UseCase,
	//publisher *broker.RabbitMQPublisher,
) *Ds18b20Controller {
	createHandler := controllers.NewCreateDs18b20Controller(createUseCase,/* publisher*/)
	viewHandler := controllers.NewGetAllDs18b20Controller(viewUseCase)

	return &Ds18b20Controller{
		createProductUseCase:   createHandler,
		viewAllProductsUseCase: viewHandler,
	}
}
func (pc *Ds18b20Controller) CreateProduct(c *gin.Context) {
	pc.createProductUseCase.SaveDs18b20Data(c)
}

func (pc *Ds18b20Controller) GetAllProducts(c *gin.Context) {
	pc.viewAllProductsUseCase.GetDs18b20Data(c)
}
