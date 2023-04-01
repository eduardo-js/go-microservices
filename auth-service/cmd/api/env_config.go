package main

import "os"

type IEnv struct {
	PORT string
	DSN  string
}

var env = IEnv{
	PORT: os.Getenv("PORT"),
	DSN:  os.Getenv("DSN"),
}
