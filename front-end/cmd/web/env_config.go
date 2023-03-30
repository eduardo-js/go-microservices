package main

type IEnv struct {
	PORT string
}

var env = IEnv{
	// PORT: os.Getenv("PORT"),
	PORT: "3000",
}
