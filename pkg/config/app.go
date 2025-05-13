package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d,err := gorm.Open("mysql",os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d 
}

func GetDB() *gorm.DB{
	return db
}