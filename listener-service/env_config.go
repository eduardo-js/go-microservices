package main

import "os"

type IEnv struct {
	RABBITMQ_CONNECTION_STRING string
}

var env = IEnv{
	RABBITMQ_CONNECTION_STRING: os.Getenv("RABBITMQ_CONNECTION_STRING"),
}
