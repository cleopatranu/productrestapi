package Models

import (
	"../Config"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllProduct(b *[]Product) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewProduct(b *Product) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneProduct(b *Product, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}
