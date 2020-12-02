package controller

import (
	"wildrift-api/database"
	"wildrift-api/errors"
	"wildrift-api/model"
	"wildrift-api/schema"
	"wildrift-api/serialization"

	"github.com/gin-gonic/gin"
)

func RegisterNews(rg *gin.RouterGroup) {
	r := rg.Group("/news")
	r.GET("", GetNews)
	r.GET("/detail", GetNewsDetail)

}

// 新闻列表
func GetNews(c *gin.Context) {

	var news []serialization.News
	err := database.DB.Model(model.News{}).Select("id, title, thumb_nail, date, category, link").Scan(&news).Error
	errors.HandleError("Err Query news list", err)
	Success(c, news)
}

// 新闻详情
func GetNewsDetail(c *gin.Context) {
	var params schema.IdSchema
	err := c.ShouldBindQuery(&params)
	errors.ParamsError(c.FullPath(), err)

	var detail serialization.NewsDetail
	err = database.DB.Model(model.News{}).Where("id=?", params.ID).Select("id, title, thumb_nail, date, category, link, content, description").Scan(&detail).Error
	errors.HandleError("Err Query news detail", err)
	Success(c, detail)
}
