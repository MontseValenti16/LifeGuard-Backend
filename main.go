package main

import (
	infraestructure "LifeGuard/src/DS18B20/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	infraestructure.InitTemp(r)
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
