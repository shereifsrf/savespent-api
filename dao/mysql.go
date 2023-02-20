package dao

import (
	"fmt"
	"os"

	"github.com/shereifsrf/savespent-api/dao/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySql *gorm.DB

func connectMySQL() {
	fmt.Println("Connecting to MySQL...")

	// use gorm to connect to mysql
	// use env variables to store sensitive data
	url := fmt.Sprint(os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") +
		"@tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")/" +
		os.Getenv("MYSQL_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local")
	var err error
	MySql, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}

	MySql.AutoMigrate(&model.User{})

	fmt.Println("Connected to MySQL")
}
