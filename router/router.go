package router

import (
	"wildrift-api/controller"

	"github.com/gin-gonic/gin"
)

var WebEngine *gin.Engine

func init() {
	WebEngine = gin.New()
	WebEngine.NoRoute()
	api := WebEngine.Group("/api")
	controller.RegisterChampion(api)
	controller.RegisterItems(api)
	controller.RegisterSettings(api)
}
