package main

import (
	"fmt"
	"listener-service/event"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitMQ, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitMQ.Close()
	consumer, err := event.New(rabbitMQ)
	if err != nil {
		panic(err)
	}
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Fatalln(err)
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
