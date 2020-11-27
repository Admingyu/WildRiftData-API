package controller

import (
	"net/http"
	"wildrift-api/constant"
	"wildrift-api/database"
	"wildrift-api/errors"
	"wildrift-api/model"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": constant.SUCCESS_STATUS, "data": data})
}

func Failure(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": constant.FAILURE_STATUS, "data": data})
}

func GetUserIdByOpenID(openID string) (id int) {
	err := database.DB.Model(&model.User{}).Where("open_id=?", openID).Select("id").Scan(&id).Error
	errors.HandleError("Err query user id by open_id", err)
	return id
}
