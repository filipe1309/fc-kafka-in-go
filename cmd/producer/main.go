package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChannel := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("hello", "TESTE", producer, nil, deliveryChannel)

	// Asynchronously deliver
	go DeliveryReport(deliveryChannel)

	// Synchronous delivery
	// e := <-deliveryChannel
	// msg := e.(*kafka.Message)
	// if msg.TopicPartition.Error != nil {
	// 	fmt.Println("ERROR on sent:", msg.TopicPartition.Error)
	// } else {
	// 	fmt.Println("Message sent:", msg.TopicPartition)
	// }

	fmt.Println("After calling delivery")

	producer.Flush(1 * 1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka_kafka_1:9092",
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}

	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		Key:            key,
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	}

	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				log.Println("ERROR on sent:", ev.TopicPartition.Error)
			} else {
				log.Println("Message sent:", ev.TopicPartition)
				// Sava into database that message was sent/processed
			}
		}
	}
}
