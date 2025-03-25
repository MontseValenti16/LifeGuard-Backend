package infraestructure

import (
	"github.com/gin-gonic/gin"
)

func Mpu6050Routes(r *gin.Engine, mpu6050Service *Mpu6050Controller) {
	mpu6050Group := r.Group("/mpu6050")
	{
		mpu6050Group.POST("/", mpu6050Service.CreateProduct)
		mpu6050Group.GET("/", mpu6050Service.GetAllProducts)
	}
}
