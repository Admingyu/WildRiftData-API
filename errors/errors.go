package errors

import (
	"fmt"
	"net/http"
	"time"
	"wildrift-api/config"
)

func ParamsError(msg string, err error) {
	if err != nil {
		panic(err)
	}
}

func HandleError(msg string, err error) {
	if err != nil {
		desc := fmt.Sprintf(`## %s
				### %s
				%s`, msg, err.Error(), time.Now())
		http.Get(fmt.Sprintf("https://sc.ftqq.com/%s.send?text=WildRift助手服务器异常&desp=%s", config.SERVER_CHAN_KEY, desc))
		panic(err)
	}
}
