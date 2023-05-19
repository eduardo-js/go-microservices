package main

import "os"

type IEnv struct {
	PORT        string
	BACKEND_URL string
}

var env = IEnv{
	PORT:        os.Getenv("PORT"),
	BACKEND_URL: os.Getenv("BACKEND_URL"),
}
