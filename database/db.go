package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Connect() {
	config := Config{
		Host:     "localhost",
		User:     "root",
		Password: "",
		Db:       "restapi",
	}
	var err error
	DB, err = gorm.Open("mysql", GetConnectionString(config))
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB!")
	}
	log.Println("connection was successful!!")

}
