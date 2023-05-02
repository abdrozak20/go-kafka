package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(errors.Wrap(err, "failed connect to rabbitmq"))
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("failed get channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		fmt.Println("failed get declace queue")
	}

	err = ch.Publish("", q.Name, false, false, amqp091.Publishing{
		ContentType: "text:plain",
		Body:        []byte(os.Args[1]),
	})
	if err != nil {
		fmt.Println("failed get publish message")
	}

	fmt.Println("send Message:", os.Args[1])
}
