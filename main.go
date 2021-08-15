package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-test/controller"
)

func main()  {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.Run(":8082")
}