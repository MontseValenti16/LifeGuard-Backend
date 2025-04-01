package main

import (
	"LifeGuard/src/core/middleware"
	initDs18b20 "LifeGuard/src/ds18b20/infrastructure"
	initMax30102 "LifeGuard/src/max30102/infraestructure"
	initMpu6050 "LifeGuard/src/mpu6050/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		r1 := gin.Default()
		r1.Use(middleware.MiddlewareCORS())
		initDs18b20.InitTemp(r1)
		if err := r1.Run(":8081"); err != nil {
			panic(err)
		}
	}()

	go func() {
		r2 := gin.Default()
		r2.Use(middleware.MiddlewareCORS())
		initMpu6050.Init(r2)
		if err := r2.Run(":8082"); err != nil {
			panic(err)
		}
	}()

	go func() {
		r3 := gin.Default()
		r3.Use(middleware.MiddlewareCORS())
		initMax30102.Init(r3)
		if err := r3.Run(":8083"); err != nil {
			panic(err)
		}
	}()

	select {} // Mantiene el programa corriendo
}
