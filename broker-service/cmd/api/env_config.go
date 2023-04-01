package main

import "os"

type IEnv struct {
	PORT string
}

var env = IEnv{
	PORT: os.Getenv("PORT"),
}
