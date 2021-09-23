package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var host = ViperEnvVariable("host")
var database = ViperEnvVariable("database")
var admin = ViperEnvVariable("admin")
var password = ViperEnvVariable("password")
var suffix = "?charset=utf8mb4&parseTime=True&loc=Local"

var DNS = fmt.Sprintf("%s:%s@%s/%s%s", admin, password, host, database, suffix)

func InitialConfig() {
	fmt.Println("Configuring app")
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Conversion{})
	DB.AutoMigrate(&Result{})
}
