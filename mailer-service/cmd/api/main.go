package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Config struct {
	Mailer Mail
}

func main() {
	app := Config{
		Mailer: createMail(),
	}
	log.Printf("Starting mailer @ %s\n", env.PORT)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", env.PORT),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func createMail() Mail {
	mailerPort, _ := strconv.Atoi(env.MAIL_PORT)
	m := Mail{
		Port:       mailerPort,
		Domain:     env.DOMAIN,
		Host:       env.HOST,
		Username:   env.USERNAME,
		Password:   env.PASSWORD,
		Encryption: env.ENCRYPTION,
		FromName:   env.FROM_NAME,
		From:       env.FROM_ADDRESS,
	}
	return m
}
