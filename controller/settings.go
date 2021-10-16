package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"wildrift-api/constant"
	"wildrift-api/database"
	"wildrift-api/errors"
	"wildrift-api/model"
	"wildrift-api/schema"
	"wildrift-api/serialization"
	"wildrift-api/utils"

	// "wildrift-api/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func RegisterSettings(rg *gin.RouterGroup) {
	r := rg.Group("/settings")
	r.GET("/generate204", Generate204)
	r.POST("/code", GetInfo)
	r.POST("/login", SaveUserInfo)
	r.GET("/devlogs", GetDevLogs)
	r.GET("/about", About)
	r.POST("", Setting)
}

// 小程序code登录
func GetInfo(c *gin.Context) {
	var params schema.CodeSchema
	err := c.ShouldBindJSON(&params)
	errors.ParamsError(c, err)

	openID, sessionKey := utils.CodeToJSession(params.Code)
	if openID == "" {
		Failure(c, constant.INVALID_CODE)
		return
	}

	// 缓存sessionKey
	key := fmt.Sprintf("SessionKey:%s", openID)
	err = database.RDB.Set(key, sessionKey, 0).Err()
	errors.HandleError("Err save ession to redis", err)

	// 保存openid
	user := model.User{OpenID: openID}
	err = database.DB.Where(model.User{OpenID: openID}).FirstOrCreate(&user).Error
	errors.HandleError("Err save openID", err)

	// 保存设备信息
	device := model.Device{
		UserID:   user.ID,
		Brand:    params.DeviceInfo.Brand,
		Model:    params.DeviceInfo.Model,
		Language: params.DeviceInfo.Language,
		Platform: params.DeviceInfo.Platform,
		System:   params.DeviceInfo.System,
		Version:  params.DeviceInfo.Version,
	}
	err = database.DB.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "user_id"}}, DoUpdates: clause.AssignmentColumns([]string{"brand", "model", "language", "platform", "system", "version"})}).Create(&device).Error
	errors.HandleError("Err save deviceInfo", err)

	notice := map[string]interface{}{"home": "新增符文页功能，入口在装备页面", "settings": "战绩查询，关联，功能测试不通过，太不稳定啦"}

	Success(c, map[string]interface{}{"openID": openID, "darkTheme": user.Darktheme, "notice": notice})
}

// 小程序用户信息报存
func SaveUserInfo(c *gin.Context) {
	var params schema.UserInfoSave
	err := c.ShouldBindJSON(&params)
	errors.ParamsError(c, err)

	// 解析传过来的用户信息json
	type DataStruct struct {
		NickName  string `json:"nickName"`
		Gender    int    `json:"gender"`
		Language  string `json:"language"`
		City      string `json:"city"`
		Province  string `json:"province"`
		Country   string `json:"country"`
		AvatarURL string `json:"avatarUrl"`
	}

	var info DataStruct
	err = json.Unmarshal([]byte(params.RawData), &info)
	errors.HandleError("error unmarshal userinfo", err)
	log.Println("Info:", info)

	// 保存，更新用户信息
	user := model.User{
		OpenID:    params.OpenID,
		NickName:  info.NickName,
		Gender:    info.Gender,
		Language:  info.Language,
		City:      info.City,
		Province:  info.Province,
		Country:   info.Country,
		AvatarUrl: info.AvatarURL,
	}
	err = database.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "OpenID"}},
		DoUpdates: clause.AssignmentColumns([]string{"avatar_url", "nick_name"})}).Create(&user).Error
	errors.HandleError("Err Upsert userinfo", err)

	userID := GetUserIdByOpenID(params.OpenID)
	if userID == 0 {
		Failure(c, constant.FAILURE_STATUS)
		return
	}

	// 保存网络信息
	wiFi := model.WiFiNetWork{
		UserID:    GetUserIdByOpenID(params.OpenID),
		SSID:      params.WiFiInfo.SSID,
		BSSID:     params.WiFiInfo.BSSID,
		Signal:    params.WiFiInfo.SignalStrength,
		Frequency: params.WiFiInfo.Frequency,
	}
	err = database.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&wiFi).Error
	errors.HandleError("Err save deviceInfo", err)

	// 保存剪贴板信息
	err = database.DB.Model(&model.User{}).Where("open_id=?", params.OpenID).Select("id").Scan(&user).Error
	errors.HandleError("Err query UseInfo", err)
	if user.ID != 0 {
		clickBoard := model.ClickBoard{UserID: user.ID, Content: params.ClickBoard}
		errors.HandleError("error Save ClickBoard", database.DB.Create(&clickBoard).Error)
	}
	type Dev struct {
		DevEnv bool `json:"devEn"`
	}

	Success(c, &Dev{DevEnv: false})
}

//  保存设置
func Setting(c *gin.Context) {
	var params schema.Settings
	err := c.ShouldBindJSON(&params)
	openID := params.OpenID
	darkTheme := params.DarkTheme
	err = database.DB.Model(&model.User{}).Where("open_id=?", openID).Updates(map[string]interface{}{"darktheme": darkTheme}).Error
	errors.HandleError("Err setting darkTheme", err)
	Success(c, nil)
}

// 开发日志
func GetDevLogs(c *gin.Context) {
	var data []serialization.DevLogSer
	err := database.DB.Model(&model.DevelopLog{}).Select("title, content, date, version").Order("id desc").Scan(&data).Error
	errors.HandleError("Err  DevelopLog", err)
	Success(c, data)
}

// 关于
func About(c *gin.Context) {
	var data serialization.About
	err := database.DB.Model(&model.About{}).Select("content").Scan(&data).Error
	errors.HandleError("Err  DevelopLog", err)
	Success(c, data)
}

func Generate204(c *gin.Context) {
	Success(c, nil)
}
