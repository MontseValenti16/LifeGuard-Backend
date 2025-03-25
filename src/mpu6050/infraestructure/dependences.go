package infraestructure

import (
	"LifeGuard/src/mpu6050/application"
	"LifeGuard/src/mpu6050/infraestructure/repository/mysql"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	ps := mysql.NewMySQL()
	createProduct := application.NewCreateMpu6050UseCase(ps)
	getAllProducts := application.NewGetAllMpu6050UseCase(ps)
	//publisher := rabbitmq.RabbitConnection()
	productController := NewProductController(createProduct, getAllProducts,/* publisher*/)
	Mpu6050Routes(r, productController)
}
