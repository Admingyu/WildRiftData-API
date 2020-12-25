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
	r.GET("/tree", ItemFrom)
	r.GET("/end_point", AllItems)
}

// 列表
func ItemSearch(c *gin.Context) {
	var params schema.ItemSearchSchema
	err := c.ShouldBindQuery(&params)
	errors.ParamsError(c, err)

	data := database.DB.Model(model.Items{})
	if params.Input != "" {
		data = data.Where("items.name like ?", utils.LikeFormat(params.Input))
	}

	var items []serialization.Items
	err = data.Order("id").Scan(&items).Error
	errors.HandleError("Error scan items", err)
	Success(c, items)
}

//合成路线
func ItemFrom(c *gin.Context) {
	var param schema.IDSchema
	err := c.ShouldBindQuery(&param)
	errors.ParamsError(c, err)
	type Result struct {
		ID      int
		Name    string
		Icon    string
		ChildID int
	}

	var relations []Result

	err = database.DB.Model(&model.Items{}).Joins("left join item_froms f on f.father_id=items.id").Where("items.id=?", param.ID).Select("items.id, items.name, items.icon, f.child_id").Scan(&relations).Error
	errors.HandleError("err query items form", err)

	var itemFrom utils.Item
	for _, r := range relations {
		itemFrom.ID = r.ID
		itemFrom.Name = r.Name
		itemFrom.Icon = r.Icon
		itemFrom.Children = append(itemFrom.Children, r.ChildID)
	}
	html := utils.ItemFormTree(itemFrom, allItems)
	c.Data(200, "text/html;charset=utf-8", []byte(html))
	return
}

func AllItems(c *gin.Context) {
	var param schema.IDSchema
	err := c.ShouldBindQuery(&param)
	errors.ParamsError(c, err)

	var children serialization.ItemSons
	err = database.DB.Model(&model.ItemFroms{}).Joins("join items i on i.id=item_froms.child_id").Where("item_froms.father_id=?", param.ID).Select("i.id, i.name, i.icon").Scan(&children).Error
	errors.HandleError("err query items form", err)

	//这里，最多有三个子节点，最坏的情况查三次，性能不会太低
	//加入判断，排除数据出现问题
	var qc int
	for i, c := range children {
		if qc > 4 {
			return
		}
		err = database.DB.Model(&model.ItemFroms{}).Joins("join items i on i.id=item_froms.child_id").Where("item_froms.father_id=?", c.ID).Select("i.id, i.name, i.icon").Scan(&children[i].Children).Error
		errors.HandleError("err query items form", err)
		children[i].ChildCount = len(children[i].Children)
		qc++
	}
	Success(c, children)
}
