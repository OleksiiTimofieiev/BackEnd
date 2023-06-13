package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	topic := "restart"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "admin",
		"auto.offset.reset": "smallest"})

	if err != nil {
		log.Panic(err)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Panic(err)
	}

	for {
		ev := consumer.Poll(6000)

		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("Message from the queue: %s\n", string(e.Value))
		case *kafka.Error:
			fmt.Printf("%v\n", e)
		}
	}
}
