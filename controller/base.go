package controller

import (
	"log"
	"net/http"
	"wildrift-api/constant"
	"wildrift-api/database"
	"wildrift-api/errors"
	"wildrift-api/model"
	"wildrift-api/utils"

	"github.com/gin-gonic/gin"
)

var allItems []utils.Item

func NotFound(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": constant.SUCCESS_STATUS, "data": nil})
}

func ErrorMiddle(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": constant.FAILURE_STATUS, "data": nil})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": constant.SUCCESS_STATUS, "data": data})
}

func Failure(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": constant.FAILURE_STATUS, "data": data})
}

func GetUserIdByOpenID(openID string) (id int) {
	log.Println("Getting userID of openID:", openID)
	err := database.DB.Model(&model.User{}).Where("open_id=?", openID).Select("id").Scan(&id).Error
	errors.HandleError("Err query user id by open_id", err)
	return id
}

func RetrieveAllItems() []utils.Item {
	type Relations struct {
		FatherID int
		ChildID  int
	}

	var items []model.Items
	var relations []Relations

	err := database.DB.Model(model.Items{}).Select("id, name, icon").Scan(&items).Error
	errors.HandleError("Err get all items", err)
	err = database.DB.Model(model.ItemFroms{}).Select("father_id, child_id").Scan(&relations).Error
	errors.HandleError("Err get all items relations", err)

	itemsWithChildren := make([]utils.Item, len(items))
	for i, item := range items {
		itemsWithChildren[i].ID = item.ID
		itemsWithChildren[i].Icon = item.Icon
		itemsWithChildren[i].Name = item.Name
		for _, r := range relations {
			if item.ID == r.FatherID {
				itemsWithChildren[i].Children = append(itemsWithChildren[i].Children, r.ChildID)
			}
		}
	}
	return itemsWithChildren

}

func init() {
	allItems = RetrieveAllItems()
}
