package errors

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"
	"wildrift-api/config"
	"wildrift-api/constant"

	"github.com/gin-gonic/gin"
)

type ParamsErr struct {
	Error string
}

func ParamsError(c *gin.Context, err error) {
	if err != nil {
		e := ParamsErr{Error: err.Error()}
		panic(e)
	}
	c.Next()
}

func SendServerChan(msg string, err interface{}) {
	desc := fmt.Sprintf(`## %s
						### %s
						%s`, msg, err, time.Now())
	http.Get(fmt.Sprintf("https://sc.ftqq.com/%s.send?text=WildRift助手服务器异常&desp=%s", config.SERVER_CHAN_KEY, desc))
}

func PanicRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				var errorType string
				var statusCode int
				_, match := err.(ParamsErr)
				if match {
					errorType = constant.PARAMS_ERR
					statusCode = 200
				} else {
					errorType = constant.SERVER_ERR
					statusCode = 500
					if gin.EnvGinMode == gin.ReleaseMode {
						SendServerChan("", err)
					}
				}
				c.AbortWithStatusJSON(statusCode, gin.H{"status": constant.FAILURE_STATUS, "message": errorType, "data": nil})
				log.Println(string(debug.Stack()))
			}
		}()
		c.Next()

	}
}

func HandleError(msg string, err error) {
	if err != nil {
		panic(err)
	}
}
