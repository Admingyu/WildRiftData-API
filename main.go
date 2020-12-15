package main

import (
	"fmt"
	"net/http"
	"os"
	"wildrift-api/config"
	"wildrift-api/constant"
	"wildrift-api/router"
)

func main() {
	os.Setenv("GIN_MODE", config.MODE)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", constant.SERVER_PORT),
		Handler: router.WebEngine,
	}

	server.ListenAndServe()
}
