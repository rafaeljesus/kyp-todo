package config

import (
	"encoding/json"
	nsq "github.com/nsqio/go-nsq"
	"log"
	"os"
)

var KYP_NSQ_LOOKUPD_URL = os.Getenv("KYP_NSQ_LOOKUPD_URL")

type EventBus struct {
	Producer *nsq.Producer
	Config   *nsq.Config
}

func NewEventBus(nsqUrl string) (*EventBus, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(nsqUrl, config)
	if err != nil {
		return nil, err
	}

	return &EventBus{producer, config}, nil
}

func (bus *EventBus) Emit(eventName string, message interface{}) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return bus.Producer.Publish(eventName, payload)
}

func (bus *EventBus) On(eventName string, handler func(message []byte)) {
	consumer, err := nsq.NewConsumer(eventName, "event_bus", bus.Config)
	if err != nil {
		log.Panic(err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		handler(message.Body)
		return nil
	}))

	err = consumer.ConnectToNSQLookupd(KYP_NSQ_LOOKUPD_URL)
	if err != nil {
		log.Panic(err)
	}
}
