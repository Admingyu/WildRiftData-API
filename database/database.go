package database

import (
	"fmt"
	. "wildrift-api/config"
	"wildrift-api/errors"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

func init() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_ADDR, DB_PORT, DB_DATABASE)
	conn := mysql.Open(dataSource)
	config := gorm.Config{}
	var err error
	DB, err = gorm.Open(conn, &config)
	errors.HandleError("Error connect database", err)

	//自动迁移
	err = DB.AutoMigrate(
	// model.About{},
	// model.DevelopLog{},
	// model.User{},
	// model.Device{},
	// model.ClickBoard{},
	// model.WiFiNetWork{},
	// model.News{},
	// model.NewsCategory{},
	// model.Champion{},
	// model.Role{},
	// model.ChampionRole{},
	// model.ChampionAbilities{},
	// model.ChampionSkins{},
	// model.Items{},
	// model.ItemFroms{},
	// model.ItemTypes{},
	)
	errors.HandleError("Error Migrate database", err)

	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", RDS_ADDR, RDS_PORT),
		Password: RDS_PASSWORD,
		DB:       RDS_DB,
	})

}
