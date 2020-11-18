package controller

import (
	"log"
	"wildrift-api/database"
	"wildrift-api/errors"
	"wildrift-api/model"
	"wildrift-api/schema"
	"wildrift-api/serialization"
	"wildrift-api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterChampion(rg *gin.RouterGroup) {
	r := rg.Group("/champions")
	r.GET("", ChampionSearch)
	r.GET("/detail", ChampionDetail)
}

// 列表
func ChampionSearch(c *gin.Context) {
	var params schema.ChampionSearchSchema
	err := c.ShouldBindQuery(&params)
	errors.ParamsError("params", err)

	data := database.DB.Model(model.Champion{})
	if params.Role != "" {
		data = data.Joins("left join champion_roles cr on cr.champion_id=champions.id").Joins("join roles r on r.id=cr.role_id").Where("r.machine_name=?", params.Role)
	}
	if params.DifficultyLevel != "" {
		data = data.Where("difficulty_level=?", params.DifficultyLevel)
	}
	if params.Input != "" {
		data = data.Where("champions.name like ? or champions.title like ? or name_en like ?", utils.LikeFormat(params.Input), utils.LikeFormat(params.Input), utils.LikeFormat(params.Input))
		log.Println(data)
	}

	var ser []serialization.ChampionSearch
	err = data.Select("champions.id, champions.name, champions.title, champions.image").Scan(&ser).Error
	errors.HandleError("Err Scan ChampionSearch", err)

	Success(c, ser)
	return
}

// 详情
func ChampionDetail(c *gin.Context) {
	var params schema.IdSchema
	err := c.ShouldBindQuery(&params)
	errors.ParamsError("params", err)
	log.Println(params)

	var info serialization.ChampionInfo
	data := database.DB.Model(model.Champion{}).Where("champions.id=?", params.ID).Select("champions.id, champions.name, champions.title, champions.hero_image, champions.difficulty_level, champions.video")
	err = data.Scan(&info).Error
	errors.HandleError("Err Scan ChampionInfo", err)

	var cabs []serialization.AbilityInfo
	data = database.DB.Model(model.ChampionAbilities{}).Where("champion_id=?", params.ID).Select("type, title, thumbnail, description, video, poster_image")
	err = data.Scan(&cabs).Error
	errors.HandleError("Err Scan cabs", err)

	var roles []serialization.Role
	data = database.DB.Model(model.ChampionRole{}).Joins("join roles r on r.id=champion_roles.role_id").Where("champion_roles.champion_id=?", params.ID).Select(" r.name, r.machine_name, r.icon")
	err = data.Scan(&roles).Error
	errors.HandleError("Err Scan roles", err)

	resp := gin.H{"info": info, "cabs": cabs, "roles": roles}
	Success(c, resp)
	return
}
