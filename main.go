package main

import (
	"fmt"
	"log"
	"net/http"
	"wildrift-api/constant"
	"wildrift-api/router"
)

func main() {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", constant.SERVER_PORT),
		Handler: router.WebEngine,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
