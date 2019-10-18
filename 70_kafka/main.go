package main

import (
	"fmt"
	// "strings"

	"github.com/Shopify/sarama"
	// sc "github.com/bsm/sarama-cluster"
	"os"
	"sync"
)

// run as go bulid && KAFKA=90.90.60.90:9000 TOPIC=my.cool.topic main
func main() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{os.Getenv("KAFKA")}, config)
	if err != nil {
		panic(err)
	}

	topic := os.Getenv("TOPIC")

	var wg sync.WaitGroup

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}

	resultCh := make(chan *sarama.ConsumerMessage, 1000)
	for _, partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				resultCh <- msg
				//if strings.Contains(string(msg.Value), "greenlit") {
				//resultCh <- msg
				//}
			}
		}(pc)
	}

	for msg := range resultCh {
		fmt.Println(string(msg.Value))
	}

	wg.Wait()
	defer consumer.Close()
}
