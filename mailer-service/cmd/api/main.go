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

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", env.PORT),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Mailer service running on port", env.PORT)
}

func createMail() Mail {
	mailerPort, _ := strconv.Atoi(env.MAIL_PORT)
	m := Mail{
		Port:        mailerPort,
		Domain:      env.DOMAIN,
		Host:        env.HOST,
		Username:    env.USERNAME,
		Password:    env.PASSWORD,
		Encryption:  env.ENCRYPTION,
		FromName:    env.FROM_NAME,
		FromAddress: env.FROM_ADDRESS,
	}
	return m
}
