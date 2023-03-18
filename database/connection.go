package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// db = db_trial user = root password = DanTE2401
	database, err := gorm.Open(mysql.Open("root:DanTE2401@/db_trial?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = database
	fmt.Println("Connection Opened to Database")
}
