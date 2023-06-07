package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// TODO: 43:24

type OrderPlacer struct {
	producer   *kafka.Producer
	topic      string
	deliverych chan kafka.Event
}

func NewOrderPlacer(p *kafka.Producer, topic string) *OrderPlacer {
	return &OrderPlacer{
		producer:   p,
		topic:      topic,
		deliverych: make(chan kafka.Event, 10000),
	}
}

func (op *OrderPlacer) placeOrder(orderType string, size int) error {
	var (
		format  = fmt.Sprintf("%s - %d", orderType, size)
		payload = []byte(format)
	)

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload,
	},
		op.deliverych,
	)

	if err != nil {
		log.Panic(err)
	}

	// e := <-delivery_chan
	<-op.deliverych
	return nil
}

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

	op := NewOrderPlacer(p, "restart")

	for i := 0; i < 1000; i++ {
		if err := op.placeOrder("market", i+1); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 3)
	}
	// delivery_chan := make(chan kafka.Event, 10000)
	// // TODO: enum or sort of
	// for {
	// 	err = p.Produce(&kafka.Message{
	// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 		Value:          []byte("restart"),
	// 	},
	// 		delivery_chan,
	// 	)

	// 	if err != nil {
	// 		log.Panic(err)
	// 	}

	// 	// e := <-delivery_chan
	// 	<-delivery_chan
	// 	time.Sleep(time.Second * 5)
	// }

	// // fmt.Printf("%+v\n", e.String())

	// // fmt.Printf("%+v\n", p)
}
