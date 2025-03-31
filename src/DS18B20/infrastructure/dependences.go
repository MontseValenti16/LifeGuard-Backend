package infraestructure

import (
	"LifeGuard/src/ds18b20/application"
	"LifeGuard/src/ds18b20/infrastructure/repository/mysql"
	"LifeGuard/src/ds18b20/infrastructure/routes"
	controller "LifeGuard/src/ds18b20/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	ps := mysql.NewMySQL()
	createProduct := application.NewCreateds18b20UseCase(ps)
	getAllProducts := application.NewGetAllDs18b20UseCase(ps)
	//publisher := rabbitmq.RabbitConnection()
	productController := controller.NewProductController(createProduct, getAllProducts,/* publisher*/)
	infraestructure.Ds18b20Routes(r, productController)
}
