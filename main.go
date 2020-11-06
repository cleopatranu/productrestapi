package main

import (
	"fmt"

	"./Config"
	"./Models"
	"./Routers"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", "root:@tcp(localhost:3307)/productrestapi?charset=utf8&parseTime=True&loc=Local") //Порт 3307
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})
	r := Routers.SetupRouter()
	r.Run()
}
