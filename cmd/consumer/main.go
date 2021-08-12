package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka_kafka_1:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest",
	}
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("Error consumer", err.Error())
	}
	topics := []string{"TESTE"}
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println("Message: ", string(msg.Value), msg.TopicPartition)
		}
	}
}
