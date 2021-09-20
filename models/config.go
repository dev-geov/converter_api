package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var host = ViperEnvVariable("host")
var admin = ViperEnvVariable("admin")
var password = ViperEnvVariable("password")

var DNS = fmt.Sprintf("%s:%s@%s", admin, password, host)

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
