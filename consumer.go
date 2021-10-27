package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"testQueue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("received messages: %s\n", d.Body)
		}
	}()

	fmt.Println("successfully connected to RabbitMQ instance")
	fmt.Println("[*] - waiting for messages")
	<-forever
}
