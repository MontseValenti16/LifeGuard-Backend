package infraestructure

import (
	"LifeGuard/src/max30102/application"
	controllers "LifeGuard/src/max30102/infraestructure/handlers"

	"github.com/gin-gonic/gin"
)

type Max30102Controller struct {
	createProductUseCase   *controllers.CreateMax30102Controller
	viewAllProductsUseCase *controllers.GetAllMax30102Controller
}

func NewProductController(
	createUseCase *application.CreateMax30102UseCase,
	viewUseCase *application.GetAll30102UseCase,
	//publisher *broker.RabbitMQPublisher,
) *Max30102Controller {
	createHandler := controllers.NewMax30102Controller(createUseCase,/* publisher*/)
	viewHandler := controllers.GetAll30102Controller(viewUseCase)

	return &Max30102Controller{
		createProductUseCase:   createHandler,
		viewAllProductsUseCase: viewHandler,
	}
}
func (pc *Max30102Controller) CreateProduct(c *gin.Context) {
	pc.createProductUseCase.SaveMax30102Data(c)
}

func (pc *Max30102Controller) GetAllProducts(c *gin.Context) {
	pc.viewAllProductsUseCase.GetMax30102Data(c)
}
