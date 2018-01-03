package main

import (
        "fmt"
        "github.com/confluentinc/confluent-kafka-go/kafka"
        "os"
        "encoding/json"
)

var producer *kafka.Producer
var topic, broker string

func InitKafka() (err error) {
	topic = "test2"
	broker = "10.148.0.4:9092"
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err!=nil{
		os.Exit(1)
	}
	return
}

func produceKafka(topic string, out Las_status) {
        value, err := json.Marshal(out)
	if err == nil {
		deliveryChan := make(chan kafka.Event)
		err = producer.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(value)}, deliveryChan)

		e := <-deliveryChan
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		}
		close(deliveryChan)
	}
}

