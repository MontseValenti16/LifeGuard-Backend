package main

import (
	"LifeGuard/src/core/middleware"
	"LifeGuard/src/max30102/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.MiddlewareCORS())
	infraestructure.Init(r)
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
