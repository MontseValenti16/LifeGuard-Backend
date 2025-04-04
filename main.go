package main

import (
	"LifeGuard/src/core/middleware"
	initDs18b20 "LifeGuard/src/ds18b20/infrastructure"
	initMax30102 "LifeGuard/src/max30102/infraestructure"
	initMpu6050 "LifeGuard/src/mpu6050/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.MiddlewareCORS())

	// Agrupa todas las inicializaciones en un solo router
	initDs18b20.InitTemp(r)   // Ej: /temperature
	initMpu6050.Init(r)       // Ej: /accelerometer
	initMax30102.Init(r)      // Ej: /heartrate

	// Inicia un único servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}