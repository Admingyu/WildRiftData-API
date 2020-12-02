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
	r.GET("roles", RoleList)
	r.GET("/detail", ChampionDetail)
}

// 身份列表
func RoleList(c *gin.Context) {
	var roles []serialization.Role
	err := database.DB.Model(&model.Role{}).Select("name, machine_name").Scan(&roles).Error
	errors.HandleError("Err get Role", err)
	Success(c, roles)
}

// Champion列表
func ChampionSearch(c *gin.Context) {
	var params schema.ChampionSearchSchema
	err := c.ShouldBindQuery(&params)
	errors.ParamsError("params", err)

	data := database.DB.Model(model.Champion{})
	if params.Role != "" && params.Role != "all" {
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
	data := database.DB.Model(model.Champion{}).Where("champions.id=?", params.ID).Select("champions.id, champions.name, champions.title, champions.hero_image, champions.difficulty_level, champions.video, champions.story, champions.alias_tip, champions.enemy_tip")
	err = data.Scan(&info).Error
	errors.HandleError("Err Scan ChampionInfo", err)

	var cabs []serialization.AbilityInfo
	data = database.DB.Model(model.ChampionAbilities{}).Where("champion_id=?", params.ID).Select("id, type, title, thumbnail, description, video, poster_image")
	err = data.Scan(&cabs).Error
	errors.HandleError("Err Scan cabs", err)

	var roles []serialization.Role
	data = database.DB.Model(model.ChampionRole{}).Joins("join roles r on r.id=champion_roles.role_id").Where("champion_roles.champion_id=?", params.ID).Select("r.id, r.name, r.machine_name, r.icon")
	err = data.Scan(&roles).Error
	errors.HandleError("Err Scan roles", err)

	var skins []serialization.Skins
	data = database.DB.Model(model.ChampionSkins{}).Where("champion_id=?", params.ID).Select("id, name, icon, image")
	err = data.Scan(&skins).Error
	errors.HandleError("Err Scan skins", err)

	resp := gin.H{"info": info, "spells": cabs, "roles": roles, "skins": skins}
	Success(c, resp)
	return
}
