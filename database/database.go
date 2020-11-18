package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dataSource := "root:admingyu@tcp(localhost:3306)/LOL?charset=utf8mb4&parseTime=True&loc=Local"
	conn := mysql.Open(dataSource)
	config := gorm.Config{}
	var err error
	DB, err = gorm.Open(conn, &config)
	if err != nil {
		panic(err)
	}

}
