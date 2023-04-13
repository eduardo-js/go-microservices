package main

import "os"

type IEnv struct {
	PORT                       string
	RABBITMQ_CONNECTION_STRING string
}

var env = IEnv{
	PORT:                       os.Getenv("PORT"),
	RABBITMQ_CONNECTION_STRING: os.Getenv("RABBITMQ_CONNECTION_STRING"),
}
