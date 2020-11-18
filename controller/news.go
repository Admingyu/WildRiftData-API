package controller

import "github.com/gin-gonic/gin"

func RegisterNews(rg *gin.RouterGroup) {
	r := rg.Group("/news")
	r.GET("", GetNews)

}

func GetNews(c *gin.Context) {

}
