package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	nsq "github.com/nsqio/go-nsq"
)

type BusMessage struct {
	Timestamp int64  `json:"dt"`
	UserAgent string `json:"ua"`
	Source    string `json:"source"`
}

type NoopNSQLogger struct{}

func (l *NoopNSQLogger) Output(calldepth int, s string) error {
	return nil
}

type MessageHandler struct{}

func (h *MessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return errors.New("body is blank re-enqueue message")
	}

	var data BusMessage
	if err := json.Unmarshal(m.Body, &data); err != nil {
		return err
	}
	log.Printf("%+v", data.UserAgent)

	return nil
}

func main() {
	config := nsq.NewConfig()

	consumer, err := nsq.NewConsumer("my-topic-site", "ch", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.ChangeMaxInFlight(200)
	consumer.SetLogger(
		&NoopNSQLogger{},
		nsq.LogLevelError,
	)

	consumer.AddConcurrentHandlers(
		&MessageHandler{},
		20,
	)

	if err := consumer.ConnectToNSQD("127.0.0.1:32781"); err != nil {
		log.Fatal(err)
	}

	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)

	log.Println("Ready")

	for {
		select {
		case <-consumer.StopChan:
			return
		case <-shutdown:
			consumer.Stop()
		}
	}
}
