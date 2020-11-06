package Routers

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("product", Controllers.ListProduct)
	v1.POST("product", Controllers.AddNewProduct)
	v1.GET("product/:id", Controllers.GetOneProduct)

	return r
}
