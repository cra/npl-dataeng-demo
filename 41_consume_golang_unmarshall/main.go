package main

import (
	"encoding/json"
	"log"
	"sync"

	nsq "github.com/bitly/go-nsq"
)

type BusMessage struct {
	Timestamp int64  `json:"dt"`
	UserAgent string `json:"ua"`
	Source    string `json:"source"`
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	c, _ := nsq.NewConsumer("my-topic", "ch", config)
	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		//buf := bytes.NewBuffer(message.Body)

		var data BusMessage
		err := json.Unmarshal(message.Body, &data)
		if err == nil {
			log.Printf("%+v", data)
		}

		wg.Done()
		return err
	}))
	err := c.ConnectToNSQD("127.0.0.1:32781")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()
}
