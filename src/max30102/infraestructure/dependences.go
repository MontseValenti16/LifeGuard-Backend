package infraestructure

import (
	"LifeGuard/src/max30102/application"
	"LifeGuard/src/max30102/infraestructure/repository/mysql"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	ps := mysql.NewMySQL()
	createProduct := application.NewMax30102UseCase(ps)
	getAllProducts := application.NewGetAllMax30102UseCase(ps)
	//publisher := rabbitmq.RabbitConnection()
	productController := NewProductController(createProduct, getAllProducts,/* publisher*/)
	Max30102Routes(r, productController)
}
