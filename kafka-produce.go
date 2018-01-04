package main

import (
        "fmt"
        "github.com/confluentinc/confluent-kafka-go/kafka"
        "os"
        "encoding/json"
	"syscall"
	"os/signal"
)

var producer *kafka.Producer
var consumer *kafka.Consumer
var broker string
//var topic string

func InitKafkaProducer() (err error) {
	broker = "10.148.0.4:9092"
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err!=nil{
		os.Exit(1)
	}
	return
}

func InitKafkaConsumer() (err error) {
	broker = "10.148.0.4:9092"
	group := "test2-group"
	topic := "test2"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               broker,
		"group.id":                        group,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"}})	
	if err!=nil{
		os.Exit(1)
	}
	topics := []string{topic}
	err = consumer.SubscribeTopics(topics, nil)
	if err!=nil{
		os.Exit(2)
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

func consumeKafka() (out Las_status_array) {
	
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	
	run := true
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
				json.Unmarshal(e.Value, &out)
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}
	return
}

