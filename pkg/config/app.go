package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db * gorm.DB
)

func Connect(){
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	connString := os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(connString)
	d,err := gorm.Open("mysql",connString)

	if err != nil {
		panic(err)
	}
	db = d 
}

func GetDB() *gorm.DB{
	return db
}