package main

import (
	"bytes"
	"log"
	nsq "github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	c, _ := nsq.NewConsumer("my-topic-site", "ch", config)
	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		buf := bytes.NewBuffer(message.Body)
		log.Printf("Got a message with body: %v", buf)
		return nil
	}))
	err := c.ConnectToNSQD("127.0.0.1:32781")
	if err != nil {
		log.Panic("Could not connect")
	}
}
