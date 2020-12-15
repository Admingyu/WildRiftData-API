package controller

import (
	"crypto/md5"
	"fmt"
	"time"
	"wildrift-api/constant"
	"wildrift-api/database"
	"wildrift-api/errors"
	"wildrift-api/model"
	"wildrift-api/schema"

	"github.com/gin-gonic/gin"
)

func RegisterAdmin(rg *gin.RouterGroup) {
	r := rg.Group("/admin")
	r.POST("/news", PostNews)
	r.GET("/news", AdminNewsDetail)
}

func PostNews(c *gin.Context) {
	var params schema.NewsUpdateSchema
	err := c.ShouldBindJSON(&params)
	errors.ParamsError(c, err)

	// 时间戳超范围
	timeNow := int(time.Now().Unix())
	if timeNow-params.TimeStamp > 5 || params.TimeStamp-timeNow > 5 {
		Failure(c, constant.INVALID_TIMESTAMP)
	}

	ingredient := fmt.Sprintf("%s%d", constant.ADMIN_PASSWORD, params.TimeStamp)
	hash := md5.Sum([]byte(ingredient))
	hashString := fmt.Sprintf("%x", hash)

	// key不正确
	if hashString != params.AuthKey {
		Failure(c, constant.INVALID_AUTH_KEY)
	} else {
		row := database.DB.Model(model.News{}).Where("uniq_id=?", params.NewsID).Updates(map[string]interface{}{"content": params.Content}).RowsAffected
		Success(c, map[string]interface{}{"key": fmt.Sprintf("%x", hash), "rows": row})
	}
}

func AdminNewsDetail(c *gin.Context) {
	var params schema.NewsUniqIDSchema
	err := c.ShouldBindQuery(&params)
	errors.ParamsError(c, err)

	var news model.News
	err = database.DB.Model(model.News{}).Where("uniq_id=?", params.UniqID).Select("content, link").Scan(&news).Error
	errors.HandleError("Error get admin news detail", err)

	Success(c, map[string]interface{}{"content": news.Content, "link": news.Link})
}
