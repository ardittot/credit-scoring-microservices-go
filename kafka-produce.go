package main

import (
        "fmt"
        "github.com/confluentinc/confluent-kafka-go/kafka"
        "os"
        "encoding/json"
)

var producer *kafka.Producer
var broker string
//var topic string

func InitKafka() (err error) {
	broker = "10.148.0.4:9092"
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err!=nil{
		os.Exit(1)
	}
	return
}

func (out Las_status) ProduceKafka() {
	topic := "test2"
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

