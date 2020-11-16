package main

import (
	"bytes"
	"log"
	"sync"

	nsq "github.com/bitly/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(4)

	config := nsq.NewConfig()
	c, _ := nsq.NewConsumer("my-topic", "ch", config)
	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		buf := bytes.NewBuffer(message.Body)
		log.Printf("Got a message with body: %v", buf)
		wg.Done()
		return nil
	}))
	err := c.ConnectToNSQD("127.0.0.1:32781")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}
