package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "tinashe-Inspiron-3580:9092"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", string(ev.Value))
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "my-topic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries
	p.Flush(15 * 1000)
}
