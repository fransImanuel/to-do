package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dbReadOnlyAddress := "192.168.15.18"
	user := "root"
	pass := "L0gitech"
	port := "3306"
	dbname := "oms"
	combinedDBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=False&loc=Local", user, pass, dbReadOnlyAddress, port, dbname)
	dbReadOnly, err := gorm.Open(mysql.Open(combinedDBURL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("InitDB Address Using: ", dbReadOnlyAddress)
	return dbReadOnly
}
