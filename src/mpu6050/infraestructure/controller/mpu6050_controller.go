package infraestructure

import (
	"LifeGuard/src/mpu6050/application"
	controllers "LifeGuard/src/mpu6050/infraestructure/handlers"

	"github.com/gin-gonic/gin"
)

type Mpu6050Controller struct {
	createProductUseCase   *controllers.CreateMpu6050Controller
	viewAllProductsUseCase *controllers.GetAllMpu6050Controller
}

func NewProductController(
	createUseCase *application.CreateMpu6050UseCase,
	viewUseCase *application.GetAllMpu6050UseCase,
	//publisher *broker.RabbitMQPublisher,
) *Mpu6050Controller {
	createHandler := controllers.NewMpu6050Controller(createUseCase,/* publisher*/)
	viewHandler := controllers.NewGetAllMpu6050Controller(viewUseCase)

	return &Mpu6050Controller{
		createProductUseCase:   createHandler,
		viewAllProductsUseCase: viewHandler,
	}
}
func (pc *Mpu6050Controller) CreateProduct(c *gin.Context) {
	pc.createProductUseCase.SaveMpu6050Data(c)
}

func (pc *Mpu6050Controller) GetAllProducts(c *gin.Context) {
	pc.viewAllProductsUseCase.GetMpu6050Data(c)
}
