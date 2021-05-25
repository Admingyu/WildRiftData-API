package controller

import (
	"fmt"
	"strings"
	"time"
	"wildrift-api/database"
	"wildrift-api/errors"
	"wildrift-api/model"
	"wildrift-api/schema"

	"github.com/gin-gonic/gin"
)

// 温度数据列表
func GetRealTimeTemperature(c *gin.Context) {
	var params schema.GetTemperatureSchema
	err := c.ShouldBindQuery(&params)
	errors.ParamsError("", err)

	res, err := database.RDB.Get("TEMPERATURE-" + params.Name).Result()
	errors.HandleError("", err)
	tempTime := strings.Split(res, ",")
	Success(c, map[string]interface{}{"temperature": tempTime[0], "time": tempTime[1]})
}

// 温度数据列表
func GetTemperatureHistory(c *gin.Context) {
	var data []model.Temperature
	err := database.DB.Model(model.Temperature{}).Select("id, temperature, time").Scan(&data).Error
	errors.HandleError("Err Query data list", err)
	Success(c, data)
}

// POST温度数据
func SaveTemperature(c *gin.Context) {
	var params schema.TemperatureSchema
	err := c.ShouldBindJSON(&params)
	errors.ParamsError("temp post params err", err)

	// redis中数据取出来放到数据库
	res, err := database.RDB.Get("TEMPERATURE-" + params.Name).Result()
	errors.HandleError("", err)
	tempTime := strings.Split(res, ",")
	err = database.DB.Model(model.Temperature{}).Create(model.Temperature{Temperature: tempTime[0], Name: tempTime[1]}).Error
	errors.HandleError("Err insert temp", err)

	// 保存最新数据到redis
	err = database.RDB.Set("TEMPERATURE-"+params.Name, fmt.Sprintf("%s,%s", params.Temp, time.Now()), 0).Err()
	errors.HandleError("", err)

	Success(c, nil)
}
