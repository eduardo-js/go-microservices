package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting broker @ %s\n", env.PORT)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", env.PORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
