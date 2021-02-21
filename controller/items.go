package controller

import (
	"wildrift-api/database"
	"wildrift-api/errors"
	"wildrift-api/model"
	"wildrift-api/schema"
	"wildrift-api/serialization"
	"wildrift-api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterItems(rg *gin.RouterGroup) {
	r := rg.Group("/items")
	r.GET("", ItemSearch)
}

// 列表
func ItemSearch(c *gin.Context) {
	var params schema.ItemSearchSchema
	err := c.ShouldBindQuery(&params)
	errors.ParamsError("params", err)

	data := database.DB.Model(model.Items{})
	if params.Input != "" {
		data = data.Where("items.name like ?", utils.LikeFormat(params.Input))
	}

	var items []serialization.Items
	err = data.Order("id").Scan(&items).Error
	errors.HandleError("Error scan items", err)
	Success(c, items)
}
