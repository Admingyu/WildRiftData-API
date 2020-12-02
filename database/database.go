package database

import (
	"wildrift-api/model"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

func init() {
	dataSource := "root:admingyu@tcp(localhost:3306)/LOL?charset=utf8mb4&parseTime=True&loc=Local"
	conn := mysql.Open(dataSource)
	config := gorm.Config{}
	var err error
	DB, err = gorm.Open(conn, &config)
	if err != nil {
		panic(err)
	}

	//自动迁移
	DB.AutoMigrate(model.ClickBoard{}, model.WiFiNetWork{}, model.Device{}, model.User{}, model.News{}, model.NewsCategory{})

	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

}
