package infrastructure

import (
	"LifeGuard/src/ds18b20/application"
	controller "LifeGuard/src/ds18b20/infrastructure/controller"
	"LifeGuard/src/ds18b20/infrastructure/repository/mysql"
	"LifeGuard/src/ds18b20/infrastructure/routes"
	_ "LifeGuard/src/ds18b20/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func InitTemp(r *gin.Engine) {
	ps := mysql.NewMySQL()
	createProduct := application.NewCreateds18b20UseCase(ps)
	getAllProducts := application.NewGetAllDs18b20UseCase(ps)
	//publisher := rabbitmq.RabbitConnection()
	productController := controller.NewProductController(createProduct, getAllProducts,/* publisher*/)
	routes.Ds18b20Routes(r, productController)
}
