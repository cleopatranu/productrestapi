package Controllers

import (
	"../Models"
	"../api"
	"github.com/gin-gonic/gin"
)

func ListProduct(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProduct(&product)
	if err != nil {
		api.RespondJSON(c, 404, product)
	} else {
		api.RespondJSON(c, 200, product)
	}
}

func AddNewProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.AddNewProduct(&product)
	if err != nil {
		api.RespondJSON(c, 404, product)
	} else {
		api.RespondJSON(c, 200, product)
	}
}

func GetOneProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetOneProduct(&product, id)
	if err != nil {
		api.RespondJSON(c, 404, product)
	} else {
		api.RespondJSON(c, 200, product)
	}
}
