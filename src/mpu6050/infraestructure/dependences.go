package infraestructure

import (
	"LifeGuard/src/mpu6050/application"
	"LifeGuard/src/mpu6050/infraestructure/repository/mysql"
	"LifeGuard/src/mpu6050/infraestructure/routes"
	controller "LifeGuard/src/mpu6050/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	ps := mysql.NewMySQL()
	createProduct := application.NewCreateMpu6050UseCase(ps)
	getAllProducts := application.NewGetAllMpu6050UseCase(ps)
	//publisher := rabbitmq.RabbitConnection()
	productController := controller.NewProductController(createProduct, getAllProducts,/* publisher*/)
	infraestructure.Mpu6050Routes(r, productController)
}
