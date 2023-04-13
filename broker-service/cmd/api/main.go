package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	RabbitMQ *amqp.Connection
}

func main() {
	rabbitMQ, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitMQ.Close()

	app := Config{
		RabbitMQ: rabbitMQ,
	}
	log.Printf("Starting broker @ %s\n", env.PORT)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", env.PORT),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connect() (*amqp.Connection, error) {
	var err error
	for i := 0; i < 10; i++ {
		c, err := amqp.Dial(env.RABBITMQ_CONNECTION_STRING)
		if err != nil {
			fmt.Println("Failed to connect to RabbitMQ")
			time.Sleep(time.Duration(3*i) * time.Second)
		} else {
			log.Println("Connected to RabbitMQ!")
			return c, nil
		}
	}
	return nil, err
}
