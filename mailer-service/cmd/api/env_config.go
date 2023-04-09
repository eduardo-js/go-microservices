package main

import "os"

type IEnv struct {
	PORT         string
	DOMAIN       string
	HOST         string
	USERNAME     string
	PASSWORD     string
	ENCRYPTION   string
	FROM_NAME    string
	FROM_ADDRESS string
	MAIL_PORT    string
}

var env = IEnv{
	PORT:         os.Getenv("PORT"),
	MAIL_PORT:    os.Getenv("MAIL_PORT"),
	DOMAIN:       os.Getenv("MAIL_DOMAIN"),
	HOST:         os.Getenv("MAIL_HOST"),
	USERNAME:     os.Getenv("MAIL_USERNAME"),
	PASSWORD:     os.Getenv("MAIL_PASSWORD"),
	ENCRYPTION:   os.Getenv("MAIL_ENCRYPTION"),
	FROM_NAME:    os.Getenv("FROM_NAME"),
	FROM_ADDRESS: os.Getenv("FROM_ADDRESS"),
}
