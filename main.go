package main

import (
	"fmt"
	"net/http"
	"wildrift-api/constant"
	"wildrift-api/router"
)

func main() {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", constant.SERVER_PORT),
		Handler: router.WebEngine,
	}

	server.ListenAndServe()
}
