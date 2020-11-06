package Models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price string `json:"price"`
}

func (b *Product) TableName() string {
	return "product"
}
