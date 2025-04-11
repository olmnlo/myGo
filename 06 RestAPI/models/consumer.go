package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

var product Job

func ConsumeTopic(topicName string) Job {
	productChannel := make(chan Job)
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // Start from the earliest message

	client, err := sarama.NewConsumer([]string{Broker}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer client.Close()

	// Get all partitions for the topic
	partitions, err := client.Partitions(topicName)
	if err != nil {
		log.Fatalf("Error retrieving partitions: %v", err)
	}

	// Consume from all partitions
	for _, partition := range partitions {
		go func(partition int32) {
			partitionConsumer, err := client.ConsumePartition(topicName, partition, sarama.OffsetOldest)
			if err != nil {
				log.Fatalf("Error creating partition consumer: %v", err)
			}
			defer partitionConsumer.Close()

			for msg := range partitionConsumer.Messages() {
				// var product producer.Job
				err := json.Unmarshal(msg.Value, &product)
				if err != nil {
					log.Printf("Failed to parse message: %v", err)
				} else {
					// Process the product message here
					fmt.Printf("Consumed from partition %d: %+v\n", partition, product)
					productChannel <- product
				}
			}
		}(partition)
	}

	// Block here to wait for the first product (or until a timeout occurs)
	product := <-productChannel
	fmt.Printf("Consumed product: %+v\n", product)

	// Close the channel after receiving the product
	close(productChannel)

	// Return the consumed product
	return product
}
