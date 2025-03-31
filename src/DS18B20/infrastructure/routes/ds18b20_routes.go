package routes
import (
	"LifeGuard/src/ds18b20/infrastructure/controller"
	"github.com/gin-gonic/gin"
)

func Ds18b20Routes(r *gin.Engine, ds18b20Service *infraestructure.Ds18b20Controller) {
	ds18b20Group := r.Group("/ds18b20")
	{
		ds18b20Group.POST("/", ds18b20Service.CreateProduct)
		ds18b20Group.GET("/", ds18b20Service.GetAllProducts)
	}
}
