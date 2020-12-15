package router

import (
	"wildrift-api/config"
	"wildrift-api/controller"
	"wildrift-api/errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var WebEngine *gin.Engine

func init() {

	//跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{config.API_SERVER}
	corsConfig.AllowMethods = []string{"GET", "POST"}

	WebEngine = gin.New()
	WebEngine.NoRoute()
	WebEngine.Use(gin.Logger(), cors.New(corsConfig), errors.PanicRecovery())
	api := WebEngine.Group("/api")
	controller.RegisterChampion(api)
	controller.RegisterItems(api)
	controller.RegisterSettings(api)
	controller.RegisterNews(api)
	controller.RegisterAdmin(api)
}
