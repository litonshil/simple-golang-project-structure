package connection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// ConnectDB connect with mysql database
func ConnectDB() {

	dsn := "root:@tcp(127.0.0.1:3306)/dbName?charset=utf8mb4&parseTime=True&loc=Local"
	dB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db = dB
	fmt.Println("mysql connection successful...")
	//db.AutoMigrate(&Stock{})
}

// GetDB return database connection
func GetDB() *gorm.DB {
	return db
}
