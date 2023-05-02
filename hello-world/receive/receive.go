package main

import (
	"fmt"
	"log"

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

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	fmt.Println("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
