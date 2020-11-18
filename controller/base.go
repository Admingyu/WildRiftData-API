package controller

import (
	"net/http"
	"wildrift-api/constant"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": constant.SUCCESS_STATUS, "data": data})
}

func Failure(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"status": constant.FAILURE_STATUS, "data": data})
}
