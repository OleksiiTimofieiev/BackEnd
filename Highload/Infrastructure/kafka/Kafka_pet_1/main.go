package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	topic := "restart"

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "admin",
		"acks":              "all"})

	if err != nil {
		log.Panic(err)
	}

	go func() {
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
	}()

	delivery_chan := make(chan kafka.Event, 10000)
	// TODO: enum or sort of
	for {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte("restart"),
		},
			delivery_chan,
		)

		if err != nil {
			log.Panic(err)
		}

		// e := <-delivery_chan
		<-delivery_chan
		time.Sleep(time.Second * 5)
	}

	// fmt.Printf("%+v\n", e.String())

	// fmt.Printf("%+v\n", p)
}

// 32:57
